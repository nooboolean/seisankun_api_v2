package usecases

import (
	"github.com/nooboolean/seisankun_api_v2/infrastructures/models"
	"github.com/nooboolean/seisankun_api_v2/infrastructures/repositories"
	responses "github.com/nooboolean/seisankun_api_v2/presentations/responses/travels"
)

func NewTravelUsecase() *travelUsecase {
	return &travelUsecase{}
}

type travelUsecase struct{}

func (u *travelUsecase) Get(travel_key string) (*responses.TravelGetResponse, error) {
	travel := models.Travel{}
	members := []models.Member{}
	r := repositories.NewTravelRepository()
	if err := r.Find(travel_key, &travel, &members); err != nil {
		return nil, err
	}
	response := responses.TravelGetResponse{Travel: travel, Members: members}
	return &response, nil
}

func (u *travelUsecase) Register(members []*models.Member, travel *models.Travel) error {
	tr := repositories.NewTravelRepository()
	if err := tr.Store(travel); err != nil {
		return err
	}

	mr := repositories.NewMemberRepository()
	mtr := repositories.NewMemberTravelRepository()
	for _, member := range members {
		if err := mr.Store(member); err != nil {
			return err
		}
		if err := mtr.Store(int(member.ID), int(travel.ID)); err != nil {
			return err
		}
	}
	return nil
}
