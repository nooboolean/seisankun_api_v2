package responses

import (
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
)

type PaymentGetRequest struct {
	Payment Payment `json:"payment" binding:"required,dive"`
}

type Payment struct {
	ID        int              `json:"id"`
	TravelKey string           `json:"travel_key"`
	PayerId   int              `json:"payer_id"`
	Borrowers domain.Borrowers `json:"borrowers"`
	Title     string           `json:"title"`
	Amount    int              `json:"amount"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
