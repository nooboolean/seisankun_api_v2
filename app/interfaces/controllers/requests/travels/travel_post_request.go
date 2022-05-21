package requests

import "github.com/nooboolean/seisankun_api_v2/domain"

type TravelPostRequest struct {
	Members domain.Members `json:"members" binding:"required,dive"`
	Travel  domain.Travel  `json:"travel" binding:"required,dive"`
}
