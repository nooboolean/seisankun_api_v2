package responses

import (
	"github.com/nooboolean/seisankun_api_v2/infrastructures/models"
)

type TravelGetResponse struct {
	Travel  models.Travel   `json:"travel"`
	Members []models.Member `json:"members"`
}
