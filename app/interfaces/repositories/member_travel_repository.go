package repositories

import (
	"database/sql"

	"github.com/nooboolean/seisankun_api_v2/domain"
)

type MemberTravelRepository struct {
	SqlHandler
}

func (r *MemberTravelRepository) FindByTravelKey(travel_key string) (memberTravelList domain.MemberTravelList, err error) {
	if err = r.Model(&domain.MemberTravel{}).Joins(`
																								LEFT JOIN travels ON member_travel.travel_id = travels.id
																				`).Where("travels.travel_key = @travel_key", sql.Named("travel_key", travel_key)).Scan(&memberTravelList).Error; err != nil {
		return
	}
	return
}

func (r *MemberTravelRepository) FindByMemberId(member_id int) (memberTravel domain.MemberTravel, err error) {
	if err = r.Model(&domain.MemberTravel{}).Where("member_id = @member_id", sql.Named("member_id", member_id)).Scan(&memberTravel).Error; err != nil {
		return
	}
	return
}

func (r *MemberTravelRepository) StoreList(memberTravelList domain.MemberTravelList) error {
	return r.Create(&memberTravelList).Error
}

func (r *MemberTravelRepository) Store(memberTravel domain.MemberTravel) error {
	return r.Create(&memberTravel).Error
}

func (r *MemberTravelRepository) DeleteList(memberTravelList domain.MemberTravelList) error {
	memberTravelListIds := []int{}
	for _, memberTravel := range memberTravelList {
		memberTravelListIds = append(memberTravelListIds, int(memberTravel.ID))
	}
	return r.Where("id IN (?)", memberTravelListIds).Delete(&domain.MemberTravel{}).Error
}

func (r *MemberTravelRepository) Delete(member_travel domain.MemberTravel) error {
	return r.Model(&member_travel).Delete(&member_travel).Error
}
