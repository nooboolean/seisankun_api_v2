package repositories

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
)

type BorrowMoneyRepository struct {
	SqlHandler
}

func (r *BorrowMoneyRepository) FindByPaymentId(payment_id int) (borrow_money_list domain.BorrowMoneyList, err error) {
	if err = r.Model(&domain.BorrowMoney{}).Where("payment_id = ?", payment_id).Scan(&borrow_money_list).Error; err != nil {
		return
	}
	return
}

func (r *BorrowMoneyRepository) FindByMemberIdWithPayment(member_id int) (borrow_money_list domain.BorrowMoneyList, err error) {
	if err = r.Model(&domain.BorrowMoney{}).Where("borrower_id = ?", member_id).Scan(&borrow_money_list).Error; err != nil {
		return
	}
	return
}

func (r *BorrowMoneyRepository) Store(borrow_money_list domain.BorrowMoneyList) (err error) {
	if err = r.Create(&borrow_money_list).Error; err != nil {
		return
	}
	return
}

func (r *BorrowMoneyRepository) Delete(payment_id int) (err error) {
	if err = r.Where("payment_id = ?", payment_id).Delete(&domain.BorrowMoney{}).Error; err != nil {
		return
	}
	return
}
