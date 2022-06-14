package controllers

import (
	"net/http"

	"github.com/nooboolean/seisankun_api_v2/domain"
	payment_history_requests "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/requests/payment_history"
	payment_requests "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/requests/payments"
	"github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/errors"
	payment_history_responses "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/payment_history"
	payment_responses "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/payments"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/usecases"

	"github.com/gin-gonic/gin"
)

type paymentController struct {
	Interactor *usecases.PaymentInteractor
}

func NewPaymentController(sqlHandler repositories.SqlHandler) *paymentController {
	return &paymentController{
		Interactor: &usecases.PaymentInteractor{
			TravelRepository: &repositories.TravelRepository{
				SqlHandler: sqlHandler,
			},
			MemberRepository: &repositories.MemberRepository{
				SqlHandler: sqlHandler,
			},
			PaymentRepository: &repositories.PaymentRepository{
				SqlHandler: sqlHandler,
			},
			BorrowMoneyRepository: &repositories.BorrowMoneyRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *paymentController) Show(c *gin.Context) {
	var request payment_requests.PaymentGetRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	travel_key, payment, borrowers, err := controller.Interactor.Get(request.PaymentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := payment_responses.PaymentGetRequest{
		Payment: payment_responses.Payment{
			ID:        int(payment.ID),
			TravelKey: travel_key,
			PayerId:   payment.PayerId,
			Borrowers: borrowers,
			Title:     payment.Title,
			Amount:    payment.Amount,
			CreatedAt: payment.CreatedAt,
			UpdatedAt: payment.UpdatedAt,
		},
	}
	c.JSON(http.StatusOK, response)
}

func (controller *paymentController) Index(c *gin.Context) {
	var request payment_history_requests.PaymentHistoryGetRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	payments, err := controller.Interactor.GetPayments(request.TravelKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response_payments := make([]payment_history_responses.Payment, 0, len(payments))
	for _, payment_member := range payments {
		response_payment := payment_history_responses.Payment{
			ID:        int(payment_member.ID),
			TravelKey: request.TravelKey,
			PayerId:   payment_member.PayerId,
			PayerName: payment_member.Member.Name,
			Title:     payment_member.Title,
			Amount:    payment_member.Amount,
			CreatedAt: payment_member.CreatedAt,
			UpdatedAt: payment_member.UpdatedAt,
		}
		response_payments = append(response_payments, response_payment)
	}
	response := payment_history_responses.PaymentHistoryGetResponse{Payments: response_payments}
	c.JSON(http.StatusOK, response)
}

func (controller *paymentController) Create(c *gin.Context) {
	var request payment_requests.PaymentPostRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	payment := domain.Payment{
		PayerId: request.Payment.PayerId,
		Title:   request.Payment.Title,
		Amount:  request.Payment.Amount,
	}

	borrow_money_list := make(domain.BorrowMoneyList, 0, len(request.Payment.Borrowers))
	for _, borrower := range request.Payment.Borrowers {
		borrow_money := domain.BorrowMoney{
			BorrowerId: borrower.BorrowerId,
		}
		borrow_money_list = append(borrow_money_list, borrow_money)
	}

	payment_id, err := controller.Interactor.Register(request.Payment.TravelKey, payment, borrow_money_list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := payment_responses.PaymentPostResponse{PaymentId: payment_id}
	c.JSON(http.StatusOK, response)
}

func (controller *paymentController) Update(c *gin.Context) {
	var request payment_requests.PaymentPutRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	payment := domain.Payment{
		ID:      uint(request.Payment.ID),
		PayerId: request.Payment.PayerId,
		Title:   request.Payment.Title,
		Amount:  request.Payment.Amount,
	}

	borrow_money_list := make(domain.BorrowMoneyList, 0, len(request.Payment.Borrowers))
	for _, borrower := range request.Payment.Borrowers {
		borrow_money := domain.BorrowMoney{
			BorrowerId: borrower.BorrowerId,
		}
		borrow_money_list = append(borrow_money_list, borrow_money)
	}

	payment_id, err := controller.Interactor.Update(request.Payment.TravelKey, payment, borrow_money_list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := payment_responses.PaymentPutResponse{PaymentId: payment_id}
	c.JSON(http.StatusOK, response)
}

func (controller *paymentController) Delete(c *gin.Context) {
	var request payment_requests.PaymentDeleteRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	err := controller.Interactor.Delete(request.PaymentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := payment_responses.PaymentDeleteResponse{Message: "success"}
	c.JSON(http.StatusOK, response)
}
