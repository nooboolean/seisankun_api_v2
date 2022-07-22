package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"gorm.io/gorm"
)

type TravelRepository struct {
	Db SqlHandler
}

func (r *TravelRepository) FindByTravelKey(travel_key string) (travel domain.Travel, err error) {
	if err = r.Db.Where(&domain.Travel{TravelKey: travel_key}).First(&travel).Error; err != nil {
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
	if err = r.Db.First(&travel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = domain.Errorf(codes.NotFound, "Failed to find travel - %s", err)
			return
		}
		err = domain.Errorf(codes.Database, "Failed to find travel  - %s", err)
		return
	}
	return
}

func (r *TravelRepository) Store(ctx context.Context, travel domain.Travel) (created_travel domain.Travel, err error) {
	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&travel).Error
	} else {
		err = r.Db.Create(&travel).Error
	}

	created_travel = travel
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create travel  - %s", err)
		return
	}
	return
}

func (r *TravelRepository) Update(t domain.Travel) (travel domain.Travel, err error) {
	if err = r.Db.Model(&t).Updates(&t).Find(&travel).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to update travel  - %s", err)
		return
	}
	return
}

func (r *TravelRepository) Delete(ctx context.Context, travel domain.Travel) (err error) {
	deleted_travel := domain.DeletedTravel{
		ID:        travel.ID,
		Name:      travel.Name,
		TravelKey: travel.TravelKey,
		CreatedAt: travel.CreatedAt,
		UpdatedAt: travel.UpdatedAt,
		DeletedAt: time.Now(),
	}

	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&deleted_travel).Error
	} else {
		err = r.Db.Create(&deleted_travel).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_travel  - %s", err)
		return
	}

	if ok {
		err = db.Model(&travel).Delete(&travel).Error
	} else {
		err = r.Db.Model(&travel).Delete(&travel).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete travel  - %s", err)
		return
	}
	return
}
