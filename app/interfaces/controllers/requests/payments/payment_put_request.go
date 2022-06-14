package requests

type PaymentPutRequest struct {
	Payment Payment `json:"payment" binding:"required,dive"`
}
