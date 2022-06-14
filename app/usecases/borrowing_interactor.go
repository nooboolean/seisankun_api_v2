package usecases

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
)

type BorrowingInteractor struct {
	MemberRepository      *repositories.MemberRepository
	BorrowMoneyRepository *repositories.BorrowMoneyRepository
}

func (i *BorrowingInteractor) Get(travel_key string) (members domain.Members, err error) {
	members, err = i.MemberRepository.FindByTravelKeyWithBorrowMoneyListAndPayments(travel_key)
	if err != nil {
		return
	}
	return
}

func (i *BorrowingInteractor) GetHistory(member_id int) (member domain.Member, borrow_money_list domain.BorrowMoneyList, err error) {
	member, err = i.MemberRepository.FindByMemberIdWithBorrowMoneyListAndPayments(member_id)
	if err != nil {
		return
	}

	borrow_money_list, err = i.BorrowMoneyRepository.FindByMemberIdWithPayment(member_id)
	if err != nil {
		return
	}
	return
}
