package domain

import (
	"time"
)

type Payment struct {
	ID              uint            `gorm:"primary_key" json:"id"`
	TravelId        int             `json:"travel_id"`
	Travel          Travel          `gorm:"foreignKey:TravelId;references:ID" json:"travel,omitempty"`
	PayerId         int             `json:"payer_id"`
	Member          Member          `gorm:"foreignKey:PayerId;references:ID" json:"member,omitempty"`
	Title           string          `json:"title"`
	Amount          int             `json:"amount"`
	BorrowMoneyList BorrowMoneyList `gorm:"foreignKey:PaymentId;references:ID" json:"borrow_money_list,omitempty"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type Payments []Payment

func (Payment) TableName() string {
	return "payments"
}
