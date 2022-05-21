package usecases

import "github.com/nooboolean/seisankun_api_v2/domain"

type MemberRepository interface {
	FindByTravelKey(string) (domain.Members, error)
	Store(domain.Members) error
}
