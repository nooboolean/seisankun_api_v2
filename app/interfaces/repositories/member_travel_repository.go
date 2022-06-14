package repositories

import (
	"database/sql"
	"time"

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

func (r *MemberTravelRepository) DeleteList(member_travel_list domain.MemberTravelList) (err error) {
	deleted_member_travel_list := make(domain.DeletedMemberTravelList, 0, len(member_travel_list))
	deleted_at := time.Now()
	for _, member_travel := range member_travel_list {
		deleted_member_travel := domain.DeletedMemberTravel{
			ID:        member_travel.ID,
			MemberId:  member_travel.MemberId,
			TravelId:  member_travel.TravelId,
			CreatedAt: member_travel.CreatedAt,
			DeletedAt: deleted_at,
		}
		deleted_member_travel_list = append(deleted_member_travel_list, deleted_member_travel)
	}
	err = r.Create(&deleted_member_travel_list).Error
	if err != nil {
		return
	}
	member_travel_list_ids := []int{}
	for _, member_travel := range member_travel_list {
		member_travel_list_ids = append(member_travel_list_ids, int(member_travel.ID))
	}
	return r.Where("id IN (?)", member_travel_list_ids).Delete(&domain.MemberTravel{}).Error
}

func (r *MemberTravelRepository) Delete(member_travel domain.MemberTravel) (err error) {
	deleted_member_travel := domain.DeletedMemberTravel{
		ID:        member_travel.ID,
		MemberId:  member_travel.MemberId,
		TravelId:  member_travel.TravelId,
		CreatedAt: member_travel.CreatedAt,
		DeletedAt: time.Now(),
	}
	err = r.Create(&deleted_member_travel).Error
	if err != nil {
		return
	}

	return r.Model(&member_travel).Delete(&member_travel).Error
}
