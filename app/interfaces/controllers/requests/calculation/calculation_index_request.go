package requests

type CalculationIndexRequest struct {
	TravelKey string `form:"travel_key" binding:"required"`
}
