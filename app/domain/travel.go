package domain

import (
	"time"
)

type Travel struct {
	ID               uint             `gorm:"primary_key" json:"id"`
	Name             string           `gorm:"not null" json:"name"`
	TravelKey        string           `gorm:"not null" json:"travel_key"`
	MemberTravelList MemberTravelList `gorm:"foreignKey:TravelId;references:ID" json:"member_travel_list,omitempty"`
	Payments         Payments         `gorm:"foreignKey:TravelId;references:ID" json:"payments,omitempty"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (Travel) TableName() string {
	return "travels"
}
