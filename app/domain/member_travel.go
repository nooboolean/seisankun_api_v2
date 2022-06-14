package domain

import (
	"time"
)

type MemberTravel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	MemberId  int       `json:"member_id"`
	Member    Member    `gorm:"foreignKey:MemberId;references:ID" json:"member,omitempty"`
	TravelId  int       `json:"travel_id"`
	Travel    Travel    `gorm:"foreignKey:TravelId;references:ID" json:"travel,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type MemberTravelList []MemberTravel

func (MemberTravel) TableName() string {
	return "member_travel"
}
