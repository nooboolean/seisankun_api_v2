package usecases

import (
	"context"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/transaction"
)

type MemberInteractor struct {
	TravelRepository       *repositories.TravelRepository
	MemberRepository       *repositories.MemberRepository
	MemberTravelRepository *repositories.MemberTravelRepository
	Transaction            transaction.Transaction
}

func (i *MemberInteractor) Register(ctx context.Context, travel_key string, member domain.Member) (member_id int, err error) {
	_, err = i.Transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		created_member, err := i.MemberRepository.StoreMember(ctx, member)
		member_id = int(created_member.ID)
		if err != nil {
			return "", err
		}
		travel, err := i.TravelRepository.FindByTravelKey(travel_key)
		if err != nil {
			return "", err
		}

		err = i.MemberTravelRepository.Store(ctx, domain.MemberTravel{TravelId: int(travel.ID), MemberId: member_id})
		if err != nil {
			return "", err
		}
		return "", err
	})

	if err != nil {
		return
	}
	return
}

func (i *MemberInteractor) Delete(ctx context.Context, member_id int) (err error) {
	_, err = i.Transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		member, err := i.MemberRepository.FindByIdWithBorrowMoneyListAndPayments(member_id)
		if err != nil {
			return "", err
		}

		if len(member.BorrowMoneyList) != 0 || len(member.Payments) != 0 {
			err = domain.Errorf(codes.InvalidRequest, "Bat Request - %s", "削除しようとしているMemberは、立て替えに関与しているため削除ができません")
			return "", err
		}

		member_travel, err := i.MemberTravelRepository.FindByMemberId(member_id)
		if err != nil {
			return "", err
		}

		err = i.MemberTravelRepository.Delete(ctx, member_travel)
		if err != nil {
			return "", err
		}

		err = i.MemberRepository.DeleteMember(ctx, member)
		if err != nil {
			return "", err
		}

		if err != nil {
			return "", err
		}
		return "", err
	})

	if err != nil {
		return
	}
	return
}
