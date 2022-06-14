package requests

type PaymentPostRequest struct {
	Payment Payment `json:"payment" binding:"required,dive"`
}
