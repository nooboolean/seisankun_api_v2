package usecases

import (
	"context"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/transaction"
)

type PaymentInteractor struct {
	TravelRepository       *repositories.TravelRepository
	MemberRepository       *repositories.MemberRepository
	MemberTravelRepository *repositories.MemberTravelRepository
	PaymentRepository      *repositories.PaymentRepository
	BorrowMoneyRepository  *repositories.BorrowMoneyRepository
	Transaction            transaction.Transaction
}

func (i *PaymentInteractor) Get(payment_id int) (travel_key string, payment domain.Payment, borrowers domain.Borrowers, err error) {
	payment, err = i.PaymentRepository.FindById(payment_id)
	if err != nil {
		return
	}

	members, err := i.MemberRepository.FindByPaymentId(int(payment.ID))
	if err != nil {
		return
	}

	borrowers = make(domain.Borrowers, 0, len(members))
	for _, member := range members {
		borrower := domain.Borrower{
			BorrowerId:   int(member.ID),
			BorrowerName: member.Name,
		}
		borrowers = append(borrowers, borrower)
	}

	travel, err := i.TravelRepository.FindById(payment.TravelId)
	if err != nil {
		return
	}
	travel_key = travel.TravelKey
	return
}

func (i *PaymentInteractor) GetPayments(travel_key string) (payments domain.Payments, err error) {
	payments, err = i.PaymentRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}
	return
}

func (i *PaymentInteractor) Register(ctx context.Context, travel_key string, payment domain.Payment, borrow_money_list domain.BorrowMoneyList) (payment_id int, err error) {
	_, err = i.Transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		travel, err := i.TravelRepository.FindByTravelKey(travel_key)
		if err != nil {
			return "", err
		}

		var paymentRelatedMemberIds []int
		paymentRelatedMemberIds = append(paymentRelatedMemberIds, payment.PayerId)
		for _, borrow_money := range borrow_money_list {
			paymentRelatedMemberIds = append(paymentRelatedMemberIds, borrow_money.BorrowerId)
		}

		for _, paymentRelatedMemberId := range paymentRelatedMemberIds {
			_, err = i.MemberTravelRepository.FindByMemberIdAndTravelId(paymentRelatedMemberId, int(travel.ID))
			if err != nil {
				if domain.ErrorCode(err) == codes.NotFound {
					err = domain.Errorf(codes.InvalidRequest, "Bat Request - %s", "立て替えた人もしくは立て替えられた人の中にTravelに参加していない人がいます")
				}
				return "", err
			}
		}

		payment.TravelId = int(travel.ID)

		created_payment, err := i.PaymentRepository.Store(ctx, payment)
		if err != nil {
			return "", err
		}
		payment_id = int(created_payment.ID)

		money := float64(payment.Amount) / float64(len(borrow_money_list))
		borrow_money_list_for_store := make(domain.BorrowMoneyList, 0, len(borrow_money_list))
		for _, borrow_money := range borrow_money_list {
			borrow_money.PaymentId = payment_id
			borrow_money.Money = money
			borrow_money_list_for_store = append(borrow_money_list_for_store, borrow_money)
		}
		err = i.BorrowMoneyRepository.Store(ctx, borrow_money_list_for_store)
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

func (i *PaymentInteractor) Update(ctx context.Context, travel_key string, payment domain.Payment, borrow_money_list domain.BorrowMoneyList) (payment_id int, err error) {
	_, err = i.Transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		_, err = i.PaymentRepository.FindById(int(payment.ID))
		if err != nil {
			return "", err
		}

		travel, err := i.TravelRepository.FindByTravelKey(travel_key)
		if err != nil {
			return "", err
		}

		var paymentRelatedMemberIds []int
		paymentRelatedMemberIds = append(paymentRelatedMemberIds, payment.PayerId)
		for _, borrow_money := range borrow_money_list {
			paymentRelatedMemberIds = append(paymentRelatedMemberIds, borrow_money.BorrowerId)
		}

		for _, paymentRelatedMemberId := range paymentRelatedMemberIds {
			_, err = i.MemberTravelRepository.FindByMemberIdAndTravelId(paymentRelatedMemberId, int(travel.ID))
			if err != nil {
				if domain.ErrorCode(err) == codes.NotFound {
					err = domain.Errorf(codes.InvalidRequest, "Bat Request - %s", "立て替えた人もしくは立て替えられた人の中にTravelに参加していない人がいます")
				}
				return "", err
			}
		}

		err = i.BorrowMoneyRepository.Delete(ctx, int(payment.ID))
		if err != nil {
			return "", err
		}

		money := float64(payment.Amount) / float64(len(borrow_money_list))
		create_borrow_money_list := make(domain.BorrowMoneyList, 0, len(borrow_money_list))
		for _, borrow_money := range borrow_money_list {
			borrow_money.PaymentId = int(payment.ID)
			borrow_money.Money = money
			create_borrow_money_list = append(create_borrow_money_list, borrow_money)
		}
		err = i.BorrowMoneyRepository.Store(ctx, create_borrow_money_list)
		if err != nil {
			return "", err
		}

		updated_payment, err := i.PaymentRepository.Update(ctx, payment)
		if err != nil {
			return "", err
		}
		payment_id = int(updated_payment.ID)
		return "", err
	})

	if err != nil {
		return
	}
	return
}

func (i *PaymentInteractor) Delete(ctx context.Context, payment_id int) (err error) {
	_, err = i.Transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		err = i.BorrowMoneyRepository.Delete(ctx, payment_id)
		if err != nil {
			return "", err
		}

		payment, err := i.PaymentRepository.FindById(payment_id)
		if err != nil {
			return "", err
		}

		err = i.PaymentRepository.Delete(ctx, payment)
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
