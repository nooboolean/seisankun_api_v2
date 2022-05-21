package requests

import "github.com/nooboolean/seisankun_api_v2/domain"

type MemberPostRequest struct {
	Travel domain.Travel `json:"travel" binding:"required,dive"`
	Member domain.Member `json:"member" binding:"required,dive"`
}
