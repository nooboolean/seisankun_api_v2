package repositories

import (
	"github.com/nooboolean/seisankun_api_v2/db"
	"github.com/nooboolean/seisankun_api_v2/infrastructures/models"
)

func NewMemberRepository() *memberRepository {
	return &memberRepository{}
}

type memberRepository struct{}

func (r *memberRepository) Store(member *models.Member) error {
	return db.DB.Create(member).Error
}
