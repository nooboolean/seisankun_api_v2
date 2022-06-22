package repositories

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
)

type BorrowMoneyRepository struct {
	SqlHandler
}

func (r *BorrowMoneyRepository) FindByPaymentId(payment_id int) (borrow_money_list domain.BorrowMoneyList, err error) {
	if err = r.Model(&domain.BorrowMoney{}).Where("payment_id = ?", payment_id).Find(&borrow_money_list).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find borrow_money  - %s", err)
		return
	}
	return
}

func (r *BorrowMoneyRepository) FindByPayments(payments domain.Payments) (borrow_money_list domain.BorrowMoneyList, err error) {
	paymentIds := []int{}
	for _, payment := range payments {
		paymentIds = append(paymentIds, int(payment.ID))
	}
	if err = r.Model(&domain.BorrowMoney{}).Joins("Payment").Where("Payment.id IN (?)", paymentIds).Find(&borrow_money_list).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find borrow_money  - %s", err)
		return
	}
	return
}

func (r *BorrowMoneyRepository) FindByMemberIdWithPayment(member_id int) (borrow_money_list domain.BorrowMoneyList, err error) {
	if err = r.Preload("Payment").Where("borrower_id = ?", member_id).Find(&borrow_money_list).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find borrow_money  - %s", err)
		return
	}
	return
}

func (r *BorrowMoneyRepository) Store(borrow_money_list domain.BorrowMoneyList) (err error) {
	if err = r.Create(&borrow_money_list).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to create borrow_money  - %s", err)
		return
	}
	return
}

func (r *BorrowMoneyRepository) Delete(payment_id int) (err error) {
	if err = r.Where("payment_id = ?", payment_id).Delete(&domain.BorrowMoney{}).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete borrow_money  - %s", err)
		return
	}
	return
}

func (r *BorrowMoneyRepository) DeleteList(borrow_money_list domain.BorrowMoneyList) (err error) {
	borrow_money_list_ids := []int{}
	for _, borrow_money := range borrow_money_list {
		borrow_money_list_ids = append(borrow_money_list_ids, int(borrow_money.ID))
	}
	if err = r.Where("id IN (?)", borrow_money_list_ids).Delete(&domain.BorrowMoney{}).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete borrow_money  - %s", err)
		return
	}
	return
}
