package repositories

import (
	"errors"
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	SqlHandler
}

func (r *PaymentRepository) FindById(id int) (payment domain.Payment, err error) {
	if err = r.First(&payment, id).Error; err != nil {
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
	if err = r.Preload("Member").Joins("Travel").Where("Travel.travel_key = ?", travel_key).Find(&payments).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) FindByMemberId(member_id int) (payments domain.Payments, err error) {
	if err = r.Where("payer_id = ?", member_id).Find(&payments).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) Store(payment domain.Payment) (created_payment domain.Member, err error) {
	if err = r.Create(&payment).Find(&created_payment).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to create payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) Update(payment domain.Payment) (updated_payment domain.Payment, err error) {
	if err = r.Model(&payment).Updates(&updated_payment).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to update payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) Delete(payment domain.Payment) (err error) {
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
	if err = r.Create(&deletedPayment).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_payment  - %s", err)
		return
	}
	if err = r.Model(&payment).Delete(&domain.Payment{}).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete payment  - %s", err)
		return
	}
	return
}

func (r *PaymentRepository) DeletePayments(payments domain.Payments) (err error) {
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
	if err = r.Create(&deletedPayments).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_payment  - %s", err)
		return
	}

	paymentIds := []int{}
	for _, payment := range payments {
		paymentIds = append(paymentIds, int(payment.ID))
	}
	if err = r.Where("id IN (?)", paymentIds).Delete(&domain.Payment{}).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete payment  - %s", err)
		return
	}
	return
}
