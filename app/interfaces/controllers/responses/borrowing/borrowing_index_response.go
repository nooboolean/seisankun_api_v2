package responses

type BorrowingIndexResponse struct {
	Histories Histories `json:"histories" binding:"required,dive"`
}

type Histories []History

type History struct {
	Lend   Lend   `json:"lend,omitempty"`
	Borrow Borrow `json:"borrow,omitempty"`
	Member Member `json:"member"`
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

type Member struct {
	Name string `json:"name"`
}
