package usecases

import "github.com/nooboolean/seisankun_api_v2/domain"

type DeletedTravelRepository interface {
	FindByTravelKey(string) (domain.Travel, error)
	Store(domain.Travel) (string, error)
}
