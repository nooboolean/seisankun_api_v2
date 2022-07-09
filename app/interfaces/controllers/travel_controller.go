package controllers

import (
	"net/http"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	requestsTravels "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/requests/travels"
	"github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/errors"
	responsesTravels "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/travels"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/transaction"
	"github.com/nooboolean/seisankun_api_v2/usecases"

	"github.com/gin-gonic/gin"
)

type travelController struct {
	Interactor *usecases.TravelInteractor
}

func NewTravelController(sqlHandler repositories.SqlHandler, transaction transaction.Transaction) *travelController {
	return &travelController{
		Interactor: &usecases.TravelInteractor{
			TravelRepository: &repositories.TravelRepository{
				Db: sqlHandler,
			},
			MemberRepository: &repositories.MemberRepository{
				Db: sqlHandler,
			},
			MemberTravelRepository: &repositories.MemberTravelRepository{
				Db: sqlHandler,
			},
			PaymentRepository: &repositories.PaymentRepository{
				Db: sqlHandler,
			},
			BorrowMoneyRepository: &repositories.BorrowMoneyRepository{
				Db: sqlHandler,
			},
			Transaction: transaction,
		},
	}
}

func (controller *travelController) Show(c *gin.Context) {
	var request requestsTravels.TravelGetRequest
	if err := c.BindQuery(&request); err != nil {
		err = domain.Errorf(codes.BadParams, "Bat Request Params - %s", err)
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	travelKey := request.TravelKey
	travel, members, err := controller.Interactor.Get(travelKey)
	if err != nil {
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	response := responsesTravels.TravelGetResponse{Travel: travel, Members: members}
	c.JSON(http.StatusOK, response)
}

func (controller *travelController) Create(c *gin.Context) {
	var request requestsTravels.TravelPostRequest
	if err := c.BindJSON(&request); err != nil {
		err = domain.Errorf(codes.BadParams, "Bat Request Params - %s", err)
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	travelKey, err := controller.Interactor.Register(c, request.Members, request.Travel)
	if err != nil {
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	response := responsesTravels.TravelPostResponse{TravelKey: travelKey}
	c.JSON(http.StatusOK, response)
}

func (controller *travelController) Update(c *gin.Context) {
	var request requestsTravels.TravelPutRequest
	if err := c.BindJSON(&request); err != nil {
		err = domain.Errorf(codes.BadParams, "Bat Request Params - %s", err)
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	travel, err := controller.Interactor.Update(request.Travel)
	if err != nil {
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	response := responsesTravels.TravelPutResponse{Travel: travel}
	c.JSON(http.StatusOK, response)
}

func (controller *travelController) Delete(c *gin.Context) {
	var request requestsTravels.TravelDeleteRequest
	if err := c.BindQuery(&request); err != nil {
		err = domain.Errorf(codes.BadParams, "Bat Request Params - %s", err)
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	err := controller.Interactor.Delete(c, request.TravelKey)
	if err != nil {
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	response := responsesTravels.TravelDeleteResponse{Message: "success"}
	c.JSON(http.StatusOK, response)
}
