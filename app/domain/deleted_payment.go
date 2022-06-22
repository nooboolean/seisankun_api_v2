package domain

import (
	"time"
)

type DeletedPayment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	TravelId  int       `json:"travel_id"`
	PayerId   int       `json:"payer_id"`
	Title     string    `json:"title"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type DeletedPayments []DeletedPayment

func (DeletedPayment) TableName() string {
	return "deleted_payments"
}
