package domain

import (
	"time"
)

type Member struct {
	ID               uint             `gorm:"primary_key" json:"id"`
	Name             string           `json:"name"`
	MemberTravelList MemberTravelList `gorm:"foreignKey:MemberId;references:ID" json:"member_travel_list,omitempty"`
	Payments         Payments         `gorm:"foreignKey:PayerId;references:ID" json:"payments,omitempty"`
	BorrowMoneyList  BorrowMoneyList  `gorm:"foreignKey:BorrowerId;references:ID" json:"borrow_money_list,omitempty"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

type Members []Member

func (Member) TableName() string {
	return "members"
}
