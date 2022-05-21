package requests

type TravelGetRequest struct {
	TravelKey string `form:"travel_key" binding:"required"`
}
