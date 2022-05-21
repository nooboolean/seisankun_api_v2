package requests

type TravelDeleteRequest struct {
	TravelKey string `form:"travel_key" binding:"required"`
}
