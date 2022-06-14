package requests

type CalculationIndexRequest struct {
	Results Results `json:"results" binding:"required"`
}

type Results []Result

type Result struct {
	BorrowerName string  `json:"borrower_name"`
	LenderName   string  `json:"lender_name"`
	BorrowMoney  float64 `json:"borrow_money"`
}
