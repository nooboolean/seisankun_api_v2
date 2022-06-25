package usecases

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
)

type PaymentInteractor struct {
	TravelRepository       *repositories.TravelRepository
	MemberRepository       *repositories.MemberRepository
	MemberTravelRepository *repositories.MemberTravelRepository
	PaymentRepository      *repositories.PaymentRepository
	BorrowMoneyRepository  *repositories.BorrowMoneyRepository
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

func (i *PaymentInteractor) Register(travel_key string, payment domain.Payment, borrow_money_list domain.BorrowMoneyList) (payment_id int, err error) {
	travel, err := i.TravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
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
			return
		}
	}

	payment.TravelId = int(travel.ID)

	created_payment, err := i.PaymentRepository.Store(payment)
	if err != nil {
		return
	}
	payment_id = int(created_payment.ID)

	money := float64(payment.Amount) / float64(len(borrow_money_list))
	borrow_money_list_for_store := make(domain.BorrowMoneyList, 0, len(borrow_money_list))
	for _, borrow_money := range borrow_money_list {
		borrow_money.PaymentId = payment_id
		borrow_money.Money = money
		borrow_money_list_for_store = append(borrow_money_list_for_store, borrow_money)
	}
	err = i.BorrowMoneyRepository.Store(borrow_money_list_for_store)
	if err != nil {
		return
	}

	return
}

func (i *PaymentInteractor) Update(travel_key string, payment domain.Payment, borrow_money_list domain.BorrowMoneyList) (payment_id int, err error) {
	_, err = i.PaymentRepository.FindById(int(payment.ID))
	if err != nil {
		return
	}

	travel, err := i.TravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
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
			return
		}
	}

	err = i.BorrowMoneyRepository.Delete(int(payment.ID))
	if err != nil {
		return
	}

	money := float64(payment.Amount) / float64(len(borrow_money_list))
	create_borrow_money_list := make(domain.BorrowMoneyList, 0, len(borrow_money_list))
	for _, borrow_money := range borrow_money_list {
		borrow_money.PaymentId = int(payment.ID)
		borrow_money.Money = money
		create_borrow_money_list = append(create_borrow_money_list, borrow_money)
	}
	err = i.BorrowMoneyRepository.Store(create_borrow_money_list)
	if err != nil {
		return
	}

	updated_payment, err := i.PaymentRepository.Update(payment)
	if err != nil {
		return
	}
	payment_id = int(updated_payment.ID)

	return
}

func (i *PaymentInteractor) Delete(payment_id int) (err error) {
	err = i.BorrowMoneyRepository.Delete(payment_id)
	if err != nil {
		return
	}

	payment, err := i.PaymentRepository.FindById(payment_id)
	if err != nil {
		return
	}

	err = i.PaymentRepository.Delete(payment)
	if err != nil {
		return
	}

	return
}
