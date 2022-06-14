package controllers

import (
	"net/http"

	requestsTravels "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/requests/travels"
	"github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/errors"
	responsesTravels "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/travels"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/usecases"

	"github.com/gin-gonic/gin"
)

type travelController struct {
	Interactor *usecases.TravelInteractor
}

func NewTravelController(sqlHandler repositories.SqlHandler) *travelController {
	return &travelController{
		Interactor: &usecases.TravelInteractor{
			TravelRepository: &repositories.TravelRepository{
				SqlHandler: sqlHandler,
			},
			MemberRepository: &repositories.MemberRepository{
				SqlHandler: sqlHandler,
			},
			MemberTravelRepository: &repositories.MemberTravelRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *travelController) Show(c *gin.Context) {
	var request requestsTravels.TravelGetRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	travelKey := request.TravelKey
	travel, members, err := controller.Interactor.Get(travelKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := responsesTravels.TravelGetResponse{Travel: travel, Members: members}
	c.JSON(http.StatusOK, response)
}

func (controller *travelController) Create(c *gin.Context) {
	var request requestsTravels.TravelPostRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	travelKey, err := controller.Interactor.Register(request.Members, request.Travel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := responsesTravels.TravelPostResponse{TravelKey: travelKey}
	c.JSON(http.StatusOK, response)
}

func (controller *travelController) Update(c *gin.Context) {
	var request requestsTravels.TravelPutRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	travel, err := controller.Interactor.Update(request.Travel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := responsesTravels.TravelPutResponse{Travel: travel}
	c.JSON(http.StatusOK, response)
}

func (controller *travelController) Delete(c *gin.Context) {
	var request requestsTravels.TravelDeleteRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.StandardErrorResponse{Error: errors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	err := controller.Interactor.Delete(request.TravelKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.StandardErrorResponse{Error: errors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := responsesTravels.TravelDeleteResponse{Message: "success"}
	c.JSON(http.StatusOK, response)
}
