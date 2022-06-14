package requests

type PaymentHistoryGetRequest struct {
	TravelKey string `form:"travel_key" binding:"required"`
}
