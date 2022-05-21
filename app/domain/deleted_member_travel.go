package domain

import (
	"time"
)

type DeletedMemberTravel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	MemberId  int       `json:"member_id"`
	TravelId  int       `json:"travel_id"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type DeletedMemberTravelList []DeletedMemberTravel

func (DeletedMemberTravel) TableName() string {
	return "deleted_member_travel"
}
