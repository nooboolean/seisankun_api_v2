package usecases

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/rs/xid"
)

type TravelInteractor struct {
	TravelRepository       *repositories.TravelRepository
	MemberRepository       *repositories.MemberRepository
	MemberTravelRepository *repositories.MemberTravelRepository
	PaymentRepository      *repositories.PaymentRepository
	BorrowMoneyRepository  *repositories.BorrowMoneyRepository
}

func (i *TravelInteractor) Get(travel_key string) (travel domain.Travel, members domain.Members, err error) {
	travel, err = i.TravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}
	members, err = i.MemberRepository.FindByTravelKey(travel_key)
	for j := 0; j < len(members); j++ {
		member, err := i.MemberRepository.FindByIdWithBorrowMoneyListAndPayments(int(members[j].ID))
		if err != nil {
			break
		}

		if len(member.BorrowMoneyList) != 0 || len(member.Payments) != 0 {
			members[j].CanDelete = false
		} else {
			members[j].CanDelete = true
		}
	}

	if err != nil {
		return
	}
	return
}

func (i *TravelInteractor) Register(members domain.Members, travel domain.Travel) (travel_key string, err error) {
	travel.TravelKey = xid.New().String()

	err = i.TravelRepository.Store(&travel)
	travel_key = travel.TravelKey
	if err != nil {
		return
	}
	err = i.MemberRepository.StoreMembers(members)
	if err != nil {
		return
	}

	member_travel_list := make(domain.MemberTravelList, 0, len(members))
	for _, member := range members {
		member_travel := domain.MemberTravel{
			MemberId: int(member.ID),
			TravelId: int(travel.ID),
		}
		member_travel_list = append(member_travel_list, member_travel)
	}

	err = i.MemberTravelRepository.StoreList(member_travel_list)
	if err != nil {
		return
	}
	return
}

func (i *TravelInteractor) Update(t domain.Travel) (travel domain.Travel, err error) {
	_, err = i.TravelRepository.FindById(int(t.ID))
	if err != nil {
		return
	}
	travel, err = i.TravelRepository.Update(t)
	if err != nil {
		return
	}
	return
}

func (i *TravelInteractor) Delete(travel_key string) (err error) {
	travel, err := i.TravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}

	members, err := i.MemberRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}

	member_travel_list, err := i.MemberTravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}

	payments, err := i.PaymentRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}

	borrowMoneyList, err := i.BorrowMoneyRepository.FindByPayments(payments)
	if err != nil {
		return
	}

	// NOTE: borrow_moneyのpayment_idやborrower_idに外部キー制約がかかっているので、先に削除の必要あり
	if len(borrowMoneyList) != 0 {
		err = i.BorrowMoneyRepository.DeleteList(borrowMoneyList)
		if err != nil {
			return
		}
	}
	// NOTE: paymentsのtravel_idやpayer_idに外部キー制約がかかっているので、先に削除の必要あり
	if len(payments) != 0 {
		err = i.PaymentRepository.DeletePayments(payments)
		if err != nil {
			return
		}
	}
	// NOTE: member_travelのmember_idやtravel_idに外部キー制約がかかっているので、先に削除の必要あり
	err = i.MemberTravelRepository.DeleteList(member_travel_list)
	if err != nil {
		return
	}
	err = i.TravelRepository.Delete(travel)
	if err != nil {
		return
	}
	err = i.MemberRepository.DeleteMembers(members)
	if err != nil {
		return
	}
	return
}
