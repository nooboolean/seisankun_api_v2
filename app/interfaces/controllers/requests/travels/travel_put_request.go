package requests

import "github.com/nooboolean/seisankun_api_v2/domain"

type TravelPutRequest struct {
	Travel domain.Travel `json:"travel" binding:"required,dive"`
}
