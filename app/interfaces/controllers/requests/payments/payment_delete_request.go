package requests

type PaymentDeleteRequest struct {
	PaymentId int `form:"payment_id" binding:"required"`
}
