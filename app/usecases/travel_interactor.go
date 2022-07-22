package usecases

import (
	"context"
	"fmt"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/transaction"
	"github.com/rs/xid"
)

type TravelInteractor struct {
	TravelRepository       *repositories.TravelRepository
	MemberRepository       *repositories.MemberRepository
	MemberTravelRepository *repositories.MemberTravelRepository
	PaymentRepository      *repositories.PaymentRepository
	BorrowMoneyRepository  *repositories.BorrowMoneyRepository
	Transaction            transaction.Transaction
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

func (i *TravelInteractor) Register(ctx context.Context, members domain.Members, travel domain.Travel) (travel_key string, err error) {
	travel.TravelKey = xid.New().String()
	_, err = i.Transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		err = i.TravelRepository.Store(ctx, &travel)
		travel_key = travel.TravelKey
		if err != nil {
			return "", err
		}
		err = i.MemberRepository.StoreMembers(ctx, members)
		if err != nil {
			return "", err
		}

		member_travel_list := make(domain.MemberTravelList, 0, len(members))
		for _, member := range members {
			member_travel := domain.MemberTravel{
				MemberId: int(member.ID),
				TravelId: int(travel.ID),
			}
			member_travel_list = append(member_travel_list, member_travel)
		}

		test_travel, err := i.TravelRepository.FindById(int(travel.ID))
		fmt.Println("トラベル")
		fmt.Println(test_travel)
		test_members, err := i.MemberRepository.FindByTravelKey(travel.TravelKey)
		fmt.Println("メンバーズ")
		fmt.Println(test_members)
		fmt.Println("中間テーブ")
		fmt.Println(member_travel_list)
		err = i.MemberTravelRepository.StoreList(ctx, member_travel_list)
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

func (i *TravelInteractor) Delete(ctx context.Context, travel_key string) (err error) {
	_, err = i.Transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		travel, err := i.TravelRepository.FindByTravelKey(travel_key)
		if err != nil {
			return "", err
		}

		members, err := i.MemberRepository.FindByTravelKey(travel_key)
		if err != nil {
			return "", err
		}

		member_travel_list, err := i.MemberTravelRepository.FindByTravelKey(travel_key)
		if err != nil {
			return "", err
		}

		payments, err := i.PaymentRepository.FindByTravelKey(travel_key)
		if err != nil {
			return "", err
		}

		borrowMoneyList, err := i.BorrowMoneyRepository.FindByPayments(payments)
		if err != nil {
			return "", err
		}

		// NOTE: borrow_moneyのpayment_idやborrower_idに外部キー制約がかかっているので、先に削除の必要あり
		if len(borrowMoneyList) != 0 {
			err = i.BorrowMoneyRepository.DeleteList(ctx, borrowMoneyList)
			if err != nil {
				return "", err
			}
		}
		// NOTE: paymentsのtravel_idやpayer_idに外部キー制約がかかっているので、先に削除の必要あり
		if len(payments) != 0 {
			err = i.PaymentRepository.DeletePayments(ctx, payments)
			if err != nil {
				return "", err
			}
		}
		// NOTE: member_travelのmember_idやtravel_idに外部キー制約がかかっているので、先に削除の必要あり
		err = i.MemberTravelRepository.DeleteList(ctx, member_travel_list)
		if err != nil {
			return "", err
		}
		err = i.TravelRepository.Delete(ctx, travel)
		if err != nil {
			return "", err
		}
		err = i.MemberRepository.DeleteMembers(ctx, members)
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
