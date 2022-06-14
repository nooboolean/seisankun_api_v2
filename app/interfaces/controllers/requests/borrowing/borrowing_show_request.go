package requests

type BorrowingShowRequest struct {
	TravelKey string `form:"travel_key" binding:"required"`
}
