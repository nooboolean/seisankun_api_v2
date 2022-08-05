package controllers

import (
	"net/http"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	requests "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/requests/borrowing"
	responses "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/borrowing"
	"github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/errors"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/usecases"

	"github.com/gin-gonic/gin"
)

type borrowingController struct {
	Interactor *usecases.BorrowingInteractor
}

func NewBorrowingController(sqlHandler repositories.SqlHandler) *borrowingController {
	return &borrowingController{
		Interactor: &usecases.BorrowingInteractor{
			MemberRepository: &repositories.MemberRepository{
				Db: sqlHandler,
			},
			BorrowMoneyRepository: &repositories.BorrowMoneyRepository{
				Db: sqlHandler,
			},
		},
	}
}

func (controller *borrowingController) Show(c *gin.Context) {
	var request requests.BorrowingShowRequest
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

	statuses := make(responses.Statuses, 0, len(members))
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
		member.MemberTravelList = nil
		member.Payments = nil
		member.BorrowMoneyList = nil
		statuses = append(statuses, responses.Status{Member: member, LendBorrowMoney: lendBorrowMoney})
	}
	response := responses.BorrowingShowResponse{Statuses: statuses}

	c.JSON(http.StatusOK, response)
}

func (controller *borrowingController) Index(c *gin.Context) {
	var request requests.BorrowingIndexRequest
	if err := c.BindQuery(&request); err != nil {
		err = domain.Errorf(codes.BadParams, "Bat Request Params - %s", err)
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	member, borrow_money_list, err := controller.Interactor.GetHistory(request.MemberId)
	if err != nil {
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}
	response_member := responses.Member{Name: member.Name}

	histories := make(responses.Histories, 0, len(member.Payments)+len(borrow_money_list))
	for _, payment := range member.Payments {
		lend := responses.Lend{Title: payment.Title, Money: float64(payment.Amount), PaymentId: int(payment.ID)}
		histories = append(histories, responses.History{Lend: lend, Member: response_member})
	}
	for _, borrow_money := range borrow_money_list {
		borrow := responses.Borrow{Title: borrow_money.Payment.Title, Money: borrow_money.Money, PaymentId: int(borrow_money.Payment.ID)}
		histories = append(histories, responses.History{Borrow: borrow, Member: responses.Member{Name: borrow_money.Payment.Member.Name}})
	}
	response := responses.BorrowingIndexResponse{Histories: histories}

	c.JSON(http.StatusOK, response)
}
