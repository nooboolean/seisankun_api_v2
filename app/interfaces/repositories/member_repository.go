package repositories

import (
	"database/sql"

	"github.com/nooboolean/seisankun_api_v2/domain"
)

type MemberRepository struct {
	SqlHandler
}

func (r *MemberRepository) FindByTravelKey(travel_key string) (members domain.Members, err error) {
	if err = r.Model(&domain.Member{}).Joins(`
																					LEFT JOIN member_travel ON member_travel.member_id = members.id
																					LEFT JOIN travels ON member_travel.travel_id = travels.id
																	`).Where("travels.travel_key = @travel_key", sql.Named("travel_key", travel_key)).Scan(&members).Error; err != nil {
		return
	}
	return
}

func (r *MemberRepository) FindById(id int) (member domain.Member, err error) {
	if err = r.First(&member, id).Error; err != nil {
		return
	}
	return
}

func (r *MemberRepository) StoreMembers(members domain.Members) (err error) {
	if err = r.Create(&members).Error; err != nil {
		return
	}
	return
}

func (r *MemberRepository) StoreMember(member domain.Member) (created_member domain.Member, err error) {
	if err = r.Create(&member).Scan(&created_member).Error; err != nil {
		return
	}
	return
}

func (r *MemberRepository) DeleteMembers(members domain.Members) (err error) {
	memberIds := []int{}
	for _, member := range members {
		memberIds = append(memberIds, int(member.ID))
	}
	return r.Where("id IN (?)", memberIds).Delete(&domain.Member{}).Error
}

func (r *MemberRepository) DeleteMember(member domain.Member) (err error) {
	return r.Model(&member).Delete(&member).Error
}
