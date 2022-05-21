package repositories

import (
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
)

type DeletedTravelRepository struct {
	SqlHandler
}

func (r *DeletedTravelRepository) Store(travel domain.Travel) (err error) {
	deleted_travel := domain.DeletedTravel{ID: travel.ID, Name: travel.Name, TravelKey: travel.TravelKey, CreatedAt: travel.CreatedAt, UpdatedAt: travel.UpdatedAt, DeletedAt: time.Now()}
	if err = r.Create(&deleted_travel).Error; err != nil {
		return
	}
	return
}
