package usecases

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
)

type CalculationInteractor struct {
	MemberRepository      *repositories.MemberRepository
	BorrowMoneyRepository *repositories.BorrowMoneyRepository
}

func (i *CalculationInteractor) Get(travel_key string) (members domain.Members, err error) {
	members, err = i.MemberRepository.FindByTravelKeyWithBorrowMoneyListAndPayments(travel_key)
	if err != nil {
		return
	}
	return
}
