package domain

type Borrower struct {
	BorrowerId   int    `json:"borrower_id"`
	BorrowerName string `json:"borrower_name"`
}

type Borrowers []Borrower
