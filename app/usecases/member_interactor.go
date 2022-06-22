package usecases

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
)

type MemberInteractor struct {
	TravelRepository       *repositories.TravelRepository
	MemberRepository       *repositories.MemberRepository
	MemberTravelRepository *repositories.MemberTravelRepository
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
	member, err := i.MemberRepository.FindByIdWithBorrowMoneyListAndPayments(member_id)
	if err != nil {
		return
	}

	if len(member.BorrowMoneyList) != 0 || len(member.Payments) != 0 {
		err = domain.Errorf(codes.InvalidRequest, "Bat Request - %s", "削除しようとしているMemberは、立て替えに関与しているため削除ができません")
		return
	}

	member_travel, err := i.MemberTravelRepository.FindByMemberId(member_id)
	if err != nil {
		return
	}

	err = i.MemberTravelRepository.Delete(member_travel)
	if err != nil {
		return
	}

	err = i.MemberRepository.DeleteMember(member)
	if err != nil {
		return
	}

	return
}
