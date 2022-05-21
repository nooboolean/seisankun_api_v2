package domain

import (
	"time"
)

type DeletedTravel struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	TravelKey string `gorm:"not null" json:"travel_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (DeletedTravel) TableName() string {
	return "deleted_travels"
}
