package usecases

import (
	"github.com/nooboolean/seisankun_api_v2/infrastructures/models"
	"github.com/nooboolean/seisankun_api_v2/infrastructures/repositories"
)

func NewMemberUsecase() *memberUsecase {
	return &memberUsecase{}
}

type memberUsecase struct{}

func (u *memberUsecase) Register(member *models.Member) error {
	r := repositories.NewMemberRepository()
	if err := r.Store(member); err != nil {
		return err
	}
	return nil
}
