package responses

import "github.com/nooboolean/seisankun_api_v2/domain"

type TravelPutResponse struct {
	Travel domain.Travel `json:"travel"`
}
