package responses

import (
	"github.com/nooboolean/seisankun_api_v2/domain"
)

type TravelGetResponse struct {
	Travel  domain.Travel  `json:"travel"`
	Members domain.Members `json:"members"`
}
