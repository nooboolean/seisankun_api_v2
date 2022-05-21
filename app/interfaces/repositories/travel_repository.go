package repositories

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
)

type TravelRepository struct {
	SqlHandler
}

func (r *TravelRepository) FindByTravelKey(travel_key string) (travel domain.Travel, err error) {
	if err = r.Where(&domain.Travel{TravelKey: travel_key}).First(&travel).Error; err != nil {
		return
	}
	return
}

func (r *TravelRepository) Store(travel *domain.Travel) (travel_key string, err error) {
	if err = r.Create(&travel).Error; err != nil {
		return
	}
	travel_key = travel.TravelKey
	return
}

func (r *TravelRepository) Update(t domain.Travel) (travel domain.Travel, err error) {
	if err = r.Model(&t).Updates(&t).Error; err != nil {
		return
	}
	travel = t
	return
}

func (r *TravelRepository) Delete(travel_key string) (err error) {
	if err = r.Where(&domain.Travel{TravelKey: travel_key}).Delete(&domain.Travel{}).Error; err != nil {
		return
	}
	return
}
