package repositories

import (
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
)

type DeletedMemberRepository struct {
	SqlHandler
}

func (r *DeletedMemberRepository) StoreMembers(members domain.Members) (err error) {
	deleted_members := make(domain.DeletedMembers, 0, len(members))
	deleted_at := time.Now()
	for _, member := range members {
		deleted_member := domain.DeletedMember{
			ID:        member.ID,
			Name:      member.Name,
			CreatedAt: member.CreatedAt,
			UpdatedAt: member.UpdatedAt,
			DeletedAt: deleted_at,
		}
		deleted_members = append(deleted_members, deleted_member)
	}
	if err = r.Create(&deleted_members).Error; err != nil {
		return
	}
	return
}

func (r *DeletedMemberRepository) Store(member domain.Member) (err error) {
	deleted_member := domain.DeletedMember{
		ID:        member.ID,
		Name:      member.Name,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
		DeletedAt: time.Now(),
	}
	if err = r.Create(&deleted_member).Error; err != nil {
		return
	}
	return
}
