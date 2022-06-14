package responses

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
)

type BorrowingShowResponse struct {
	Statuses Statuses `json:"statuses" binding:"required,dive"`
}

type Statuses []Status

type Status struct {
	Member          domain.Member `json:"member"`
	LendBorrowMoney float64       `json:"lend_borrow_money"`
}
