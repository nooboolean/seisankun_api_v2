package repositories

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
)

type PaymentRepository struct {
	SqlHandler
}

func (r *PaymentRepository) FindById(id int) (payment domain.Payment, err error) {
	if err = r.First(&payment, id).Error; err != nil {
		return
	}
	return
}

func (r *PaymentRepository) FindByTravelKey(travel_key string) (payments domain.Payments, err error) {
	if err = r.Debug().Preload("Member").Joins("Travel").Where("Travel.travel_key = ?", travel_key).Find(&payments).Error; err != nil {
		return
	}
	return
}

func (r *PaymentRepository) Store(payment domain.Payment) (created_payment domain.Member, err error) {
	if err = r.Create(&payment).Scan(&created_payment).Error; err != nil {
		return
	}
	return
}

func (r *PaymentRepository) Update(payment domain.Payment) (updated_payment domain.Payment, err error) {
	if err = r.Model(&payment).Updates(&updated_payment).Error; err != nil {
		return
	}
	return
}

func (r *PaymentRepository) Delete(id int) (err error) {
	if err = r.First(&domain.Payment{}, id).Delete(&domain.Payment{}).Error; err != nil {
		return
	}
	return
}
