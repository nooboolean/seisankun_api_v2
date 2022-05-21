package usecases

import (
	"fmt"
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
)

type MemberInteractor struct {
	TravelRepository              *repositories.TravelRepository
	MemberRepository              *repositories.MemberRepository
	DeletedMemberRepository       *repositories.DeletedMemberRepository
	MemberTravelRepository        *repositories.MemberTravelRepository
	DeletedMemberTravelRepository *repositories.DeletedMemberTravelRepository
}

func (i *MemberInteractor) Register(travel_key string, member domain.Member) (member_id int, err error) {
	created_member, err := i.MemberRepository.StoreMember(member)
	member_id = int(created_member.ID)
	if err != nil {
		return
	}
	travel, err := i.TravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}

	err = i.MemberTravelRepository.Store(domain.MemberTravel{TravelId: int(travel.ID), MemberId: member_id})
	return
}

func (i *MemberInteractor) Delete(member_id int) (err error) {
	member, err := i.MemberRepository.FindById(member_id)
	fmt.Println("memberを出力するよ")
	fmt.Println(member)
	if err != nil {
		return
	}

	member_travel, err := i.MemberTravelRepository.FindByMemberId(member_id)
	if err != nil {
		return
	}

	deleted_member_travel := domain.DeletedMemberTravel{
		ID:        member_travel.ID,
		MemberId:  member_travel.MemberId,
		TravelId:  member_travel.TravelId,
		CreatedAt: member_travel.CreatedAt,
		DeletedAt: time.Now(),
	}

	err = i.DeletedMemberTravelRepository.Store(deleted_member_travel)
	if err != nil {
		return
	}

	err = i.MemberTravelRepository.Delete(member_travel)
	if err != nil {
		return
	}

	err = i.DeletedMemberRepository.Store(member)
	if err != nil {
		return
	}

	err = i.MemberRepository.DeleteMember(member)
	if err != nil {
		return
	}

	return
}
