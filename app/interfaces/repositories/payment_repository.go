package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	Db SqlHandler
}

func (r *PaymentRepository) FindById(id int) (payment domain.Payment, err error) {
	if err = r.Db.First(&payment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = domain.Errorf(codes.NotFound, "Failed to find payment - %s", err)
			return
		}
		err = domain.Errorf(codes.Database, "Failed to find payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) FindByTravelKey(travel_key string) (payments domain.Payments, err error) {
	if err = r.Db.Preload("Member").Joins("Travel").Where("Travel.travel_key = ?", travel_key).Find(&payments).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) FindByMemberId(member_id int) (payments domain.Payments, err error) {
	if err = r.Db.Where("payer_id = ?", member_id).Find(&payments).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) Store(ctx context.Context, payment domain.Payment) (created_payment domain.Member, err error) {
	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&payment).Scan(&created_payment).Error
	} else {
		err = r.Db.Create(&payment).Scan(&created_payment).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) Update(ctx context.Context, payment domain.Payment) (updated_payment domain.Payment, err error) {
	db, ok := GetTx(ctx)
	if ok {
		err = db.Model(&payment).Updates(&payment).Find(&updated_payment).Error
	} else {
		err = r.Db.Model(&payment).Updates(&payment).Find(&updated_payment).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to update payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) Delete(ctx context.Context, payment domain.Payment) (err error) {
	deletedPayment := domain.DeletedPayment{
		ID:        payment.ID,
		TravelId:  payment.TravelId,
		PayerId:   payment.PayerId,
		Title:     payment.Title,
		Amount:    payment.Amount,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
		DeletedAt: time.Now(),
	}

	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&deletedPayment).Error
	} else {
		err = r.Db.Create(&deletedPayment).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_payment  - %s", err)
		return
	}

	if ok {
		err = db.Model(&payment).Delete(&payment).Error
	} else {
		err = r.Db.Model(&payment).Delete(&payment).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) DeletePayments(ctx context.Context, payments domain.Payments) (err error) {
	deletedPayments := make(domain.DeletedPayments, 0, len(payments))
	deletedAt := time.Now()
	for _, payment := range payments {
		deletedPayment := domain.DeletedPayment{
			ID:        payment.ID,
			TravelId:  payment.TravelId,
			PayerId:   payment.PayerId,
			Title:     payment.Title,
			Amount:    payment.Amount,
			CreatedAt: payment.CreatedAt,
			UpdatedAt: payment.UpdatedAt,
			DeletedAt: deletedAt,
		}
		deletedPayments = append(deletedPayments, deletedPayment)
	}

	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&deletedPayments).Error
	} else {
		err = r.Db.Create(&deletedPayments).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_payment  - %s", err)
		return
	}

	paymentIds := []int{}
	for _, payment := range payments {
		paymentIds = append(paymentIds, int(payment.ID))
	}

	if ok {
		err = db.Where("id IN (?)", paymentIds).Delete(&domain.Payment{}).Error
	} else {
		err = r.Db.Where("id IN (?)", paymentIds).Delete(&domain.Payment{}).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete payment  - %s", err)
		return
	}
	return
}
