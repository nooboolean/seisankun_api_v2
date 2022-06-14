package domain

import (
	"time"
)

type BorrowMoney struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	PaymentId  int       `json:"payment_id"`
	Payment    Payment   `gorm:"foreignKey:PaymentId;references:ID" json:"payment,omitempty"`
	BorrowerId int       `json:"borrower_id"`
	Member     Member    `gorm:"foreignKey:BorrowerId;references:ID" json:"member,omitempty"`
	Money      float64   `json:"money"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type BorrowMoneyList []BorrowMoney

func (BorrowMoney) TableName() string {
	return "borrow_money"
}
