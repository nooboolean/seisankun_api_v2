package controllers

import (
	"math"
	"net/http"
	"sort"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	requests "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/requests/calculation"
	responses "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/calculation"
	"github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/errors"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/usecases"

	"github.com/gin-gonic/gin"
)

type calculationController struct {
	Interactor *usecases.CalculationInteractor
}

func NewCalculationController(sqlHandler repositories.SqlHandler) *calculationController {
	return &calculationController{
		Interactor: &usecases.CalculationInteractor{
			MemberRepository: &repositories.MemberRepository{
				SqlHandler: sqlHandler,
			},
			BorrowMoneyRepository: &repositories.BorrowMoneyRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *calculationController) Index(c *gin.Context) {
	var request requests.CalculationIndexRequest
	if err := c.BindQuery(&request); err != nil {
		err = domain.Errorf(codes.BadParams, "Bat Request Params - %s", err)
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	members, err := controller.Interactor.Get(request.TravelKey)
	if err != nil {
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	var lenders lenders
	var borrowers borrowers
	for _, member := range members {
		var lendMoney int
		var borrowMoney float64
		var lendBorrowMoney float64
		for _, payment := range member.Payments {
			lendMoney += payment.Amount
		}
		for _, borrow_money := range member.BorrowMoneyList {
			borrowMoney += borrow_money.Money
		}
		lendBorrowMoney = float64(lendMoney) - borrowMoney
		if math.Signbit(lendBorrowMoney) {
			borrowers = append(borrowers, borrower{Name: member.Name, Money: lendBorrowMoney})
		} else {
			lenders = append(lenders, lender{Name: member.Name, Money: lendBorrowMoney})
		}
	}
	sort.SliceStable(borrowers, func(i, j int) bool { return borrowers[i].Money < borrowers[j].Money }) // NOTE: 昇順(数値が低い順)
	sort.SliceStable(lenders, func(i, j int) bool { return lenders[i].Money > lenders[j].Money })       // NOTE: 降順(数値が高い順)

	results := make(responses.Results, 0, len(members))

	for i := 0; i < len(borrowers); i++ {
		for j := 0; j < len(lenders); j++ {
			if lenders[j].Money == 0 {
				continue
			}
			var result responses.Result
			compareBorrowMoney := lenders[j].Money - (-borrowers[i].Money)
			if compareBorrowMoney > 0 {
				result.BorrowerName = borrowers[i].Name
				result.LenderName = lenders[j].Name
				result.BorrowMoney = -borrowers[i].Money
				borrowers[i].Money = 0
				lenders[j].Money = compareBorrowMoney
				results = append(results, result)
				break
			} else if compareBorrowMoney < 0 {
				result.BorrowerName = borrowers[i].Name
				result.LenderName = lenders[j].Name
				result.BorrowMoney = lenders[j].Money
				borrowers[i].Money = compareBorrowMoney
				lenders[j].Money = 0
				results = append(results, result)
			} else if compareBorrowMoney == 0 {
				result.BorrowerName = borrowers[i].Name
				result.LenderName = lenders[j].Name
				result.BorrowMoney = lenders[j].Money
				borrowers[i].Money = compareBorrowMoney
				lenders[j].Money = compareBorrowMoney
				results = append(results, result)
				break
			}
		}
	}

	c.JSON(http.StatusOK, results)
}

type lenders []lender

type lender struct {
	Name  string
	Money float64
}

type borrowers []borrower

type borrower struct {
	Name  string
	Money float64
}
