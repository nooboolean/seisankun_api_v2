package controllers

import (
	"net/http"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	requests "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/requests/members"
	"github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/errors"
	responses "github.com/nooboolean/seisankun_api_v2/interfaces/controllers/responses/members"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"github.com/nooboolean/seisankun_api_v2/transaction"
	"github.com/nooboolean/seisankun_api_v2/usecases"

	"github.com/gin-gonic/gin"
)

type memberController struct {
	Interactor *usecases.MemberInteractor
}

func NewMemberController(sqlHandler repositories.SqlHandler, transaction transaction.Transaction) *memberController {
	return &memberController{
		Interactor: &usecases.MemberInteractor{
			TravelRepository: &repositories.TravelRepository{
				Db: sqlHandler,
			},
			MemberRepository: &repositories.MemberRepository{
				Db: sqlHandler,
			},
			MemberTravelRepository: &repositories.MemberTravelRepository{
				Db: sqlHandler,
			},
			Transaction: transaction,
		},
	}
}

func (controller *memberController) Create(c *gin.Context) {
	var request requests.MemberPostRequest
	if err := c.BindJSON(&request); err != nil {
		err = domain.Errorf(codes.BadParams, "Bat Request Params - %s", err)
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	member_id, err := controller.Interactor.Register(c, request.Travel.TravelKey, request.Member)
	if err != nil {
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	response := responses.MemberPostResponse{MemberId: member_id}
	c.JSON(http.StatusOK, response)
}

func (controller *memberController) Delete(c *gin.Context) {
	var request requests.MemberDeleteRequest
	if err := c.BindQuery(&request); err != nil {
		err = domain.Errorf(codes.BadParams, "Bat Request Params - %s", err)
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	err := controller.Interactor.Delete(c, request.MemberId)
	if err != nil {
		c.JSON(errors.ToHttpStatus(err), errors.StandardErrorResponse{Error: errors.Error{Message: err.Error()}})
		return
	}

	response := responses.MemberDeleteResponse{Message: "success"}
	c.JSON(http.StatusOK, response)
}
