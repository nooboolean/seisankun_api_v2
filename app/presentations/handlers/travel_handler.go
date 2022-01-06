package handlers

import (
	"math/rand"
	"net/http"

	"github.com/nooboolean/seisankun_api_v2/infrastructures/models"
	requestsTravels "github.com/nooboolean/seisankun_api_v2/presentations/requests/travels"
	responsesErrors "github.com/nooboolean/seisankun_api_v2/presentations/responses/errors"
	responsesTravels "github.com/nooboolean/seisankun_api_v2/presentations/responses/travels"
	"github.com/nooboolean/seisankun_api_v2/usecases"

	"github.com/gin-gonic/gin"
)

func NewTravelHandler() *travelHandler {
	return &travelHandler{}
}

type travelHandler struct{}

func (h *travelHandler) Get(c *gin.Context) {
	var request requestsTravels.TravelGetRequest
	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, responsesErrors.StandardErrorResponse{Error: responsesErrors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	usecase := usecases.NewTravelUsecase()
	travelKey := request.TravelKey
	result, err := usecase.Get(travelKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responsesErrors.StandardErrorResponse{Error: responsesErrors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *travelHandler) Post(c *gin.Context) {
	var request requestsTravels.TravelPostRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, responsesErrors.StandardErrorResponse{Error: responsesErrors.Error{Message: "Bad Request.", Detail: err.Error()}})
		return
	}

	usecase := usecases.NewTravelUsecase()

	members := []*models.Member{}
	for _, member := range request.Members {
		members = append(members, &models.Member{Name: member.Name})
	}
	travel := models.Travel{Name: request.Travel.Name, TravelKey: RandomString(30)}
	if err := usecase.Register(members, &travel); err != nil {
		c.JSON(http.StatusInternalServerError, responsesErrors.StandardErrorResponse{Error: responsesErrors.Error{Message: "Internal Server Error.", Detail: err.Error()}})
		return
	}

	response := responsesTravels.TravelPostResponse{TravelKey: travel.TravelKey}
	c.JSON(http.StatusOK, response)
}

//TODO: ユニークではないため、なにかのモジュールを使って対応する
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
