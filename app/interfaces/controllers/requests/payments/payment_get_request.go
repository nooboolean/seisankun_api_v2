package requests

type PaymentGetRequest struct {
	PaymentId int `form:"payment_id" binding:"required"`
}
