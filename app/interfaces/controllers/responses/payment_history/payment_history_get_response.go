package responses

import (
	"time"
)

type PaymentHistoryGetResponse struct {
	Payments []Payment `json:"payments" binding:"required,dive"`
}

type Payment struct {
	ID        int       `json:"id"`
	TravelKey string    `json:"travel_key"`
	PayerId   int       `json:"payer_id"`
	PayerName string    `json:"payer_name"`
	Title     string    `json:"title"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
