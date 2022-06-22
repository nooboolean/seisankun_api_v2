package repositories

import (
	"errors"
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"gorm.io/gorm"
)

type TravelRepository struct {
	SqlHandler
}

func (r *TravelRepository) FindByTravelKey(travel_key string) (travel domain.Travel, err error) {
	if err = r.Where(&domain.Travel{TravelKey: travel_key}).First(&travel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = domain.Errorf(codes.NotFound, "Failed to find travel - %s", err)
			return
		}
		err = domain.Errorf(codes.Database, "Failed to find travel  - %s", err)
		return
	}
	return
}

func (r *TravelRepository) FindById(id int) (travel domain.Travel, err error) {
	if err = r.First(&travel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = domain.Errorf(codes.NotFound, "Failed to find travel - %s", err)
			return
		}
		err = domain.Errorf(codes.Database, "Failed to find travel  - %s", err)
		return
	}
	return
}

func (r *TravelRepository) Store(travel *domain.Travel) (travel_key string, err error) {
	if err = r.Create(&travel).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to create travel  - %s", err)
		return
	}
	travel_key = travel.TravelKey
	return
}

func (r *TravelRepository) Update(t domain.Travel) (travel domain.Travel, err error) {
	if err = r.Model(&t).Updates(&t).Find(&travel).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to update travel  - %s", err)
		return
	}
	return
}

func (r *TravelRepository) Delete(travel domain.Travel) (err error) {
	deleted_travel := domain.DeletedTravel{
		ID:        travel.ID,
		Name:      travel.Name,
		TravelKey: travel.TravelKey,
		CreatedAt: travel.CreatedAt,
		UpdatedAt: travel.UpdatedAt,
		DeletedAt: time.Now(),
	}
	if err = r.Create(&deleted_travel).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_travel  - %s", err)
		return
	}
	if err = r.Model(&travel).Delete(&travel).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete travel  - %s", err)
		return
	}
	return
}
