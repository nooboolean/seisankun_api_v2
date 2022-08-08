package responses

type BorrowingIndexResponse struct {
	LendHistories   LendHistories   `json:"lend_histories" binding:"required,dive"`
	BorrowHistories BorrowHistories `json:"borrow_histories" binding:"required,dive"`
}

type LendHistories []LendHistory

type LendHistory struct {
	Lend  Lend  `json:"lend,omitempty"`
	Payer Payer `json:"payer"`
}

type BorrowHistories []BorrowHistory

type BorrowHistory struct {
	Borrow Borrow `json:"borrow,omitempty"`
	Payer  Payer  `json:"payer"`
}

type Lend struct {
	Title     string  `json:"title"`
	Money     float64 `json:"money"`
	PaymentId int     `json:"payment_id"`
}

type Borrow struct {
	Title     string  `json:"title"`
	Money     float64 `json:"money"`
	PaymentId int     `json:"payment_id"`
}

type Payer struct {
	Name string `json:"name"`
}
