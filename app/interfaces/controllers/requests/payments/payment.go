package requests

import "github.com/nooboolean/seisankun_api_v2/domain"

type Payment struct {
	ID        int              `json:"id"`
	TravelKey string           `json:"travel_key"`
	PayerId   int              `json:"payer_id"`
	Borrowers domain.Borrowers `json:"borrowers"`
	Title     string           `json:"title"`
	Amount    int              `json:"amount"`
}
