package repositories

import "github.com/nooboolean/seisankun_api_v2/domain"

type DeletedMemberTravelRepository struct {
	SqlHandler
}

func (r *DeletedMemberTravelRepository) StoreList(deletedMemberTravelList domain.DeletedMemberTravelList) error {
	return r.Create(&deletedMemberTravelList).Error
}

func (r *DeletedMemberTravelRepository) Store(deletedMemberTravel domain.DeletedMemberTravel) error {
	return r.Create(&deletedMemberTravel).Error
}
