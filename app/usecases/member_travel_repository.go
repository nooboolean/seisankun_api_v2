package usecases

import "github.com/nooboolean/seisankun_api_v2/domain"

type MemberTravelRepository interface {
	Store(domain.MemberTravelList) error
}
