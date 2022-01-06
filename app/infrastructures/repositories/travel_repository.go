package repositories

import (
	"github.com/nooboolean/seisankun_api_v2/db"
	"github.com/nooboolean/seisankun_api_v2/infrastructures/models"
)

func NewTravelRepository() *travelRepository {
	return &travelRepository{}
}

type travelRepository struct{}

func (r *travelRepository) Find(travel_key string, travel *models.Travel, members *[]models.Member) error {
	if err := db.DB.Where("travel_key = ?", travel_key).First(travel).Error; err != nil {
		return err
	}
	if err := db.DB.Table("members").Joins("left join member_travel on member_travel.member_id = members.id left join travels on member_travel.travel_id = travels.id").Where("travel_key = ?", travel_key).Scan(members).Error; err != nil {
		return err
	}

	return nil
}

func (r *travelRepository) Store(travel *models.Travel) error {
	return db.DB.Create(travel).Error
}
