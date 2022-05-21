package domain

import (
	"time"
)

type Travel struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	TravelKey string `gorm:"not null" json:"travel_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Travel) TableName() string {
	return "travels"
}
