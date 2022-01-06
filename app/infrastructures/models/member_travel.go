package models

import (
	"time"
)

type MemberTravel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	MemberId  int       `json:"member_id"`
	TravelId  int       `json:"travel_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (MemberTravel) TableName() string {
	return "member_travel"
}
