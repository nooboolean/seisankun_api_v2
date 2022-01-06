package repositories

import (
	"github.com/nooboolean/seisankun_api_v2/db"
	"github.com/nooboolean/seisankun_api_v2/infrastructures/models"
)

func NewMemberTravelRepository() *memberTravelRepository {
	return &memberTravelRepository{}
}

type memberTravelRepository struct{}

func (r *memberTravelRepository) Store(memberId int, travelId int) error {
	memberTravel := models.MemberTravel{MemberId: memberId, TravelId: travelId}
	return db.DB.Create(&memberTravel).Error
}
