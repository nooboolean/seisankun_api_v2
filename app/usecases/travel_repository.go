package usecases

import "github.com/nooboolean/seisankun_api_v2/domain"

type TravelRepository interface {
	FindByTravelKey(string) (domain.Travel, error)
	Store(*domain.Travel) (string, error)
	Update(domain.Travel) (domain.Travel, error)
	Delete(string) error
}
