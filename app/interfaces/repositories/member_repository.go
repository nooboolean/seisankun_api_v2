package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"gorm.io/gorm"
)

type MemberRepository struct {
	Db SqlHandler
}

func (r *MemberRepository) FindByTravelKey(travel_key string) (members domain.Members, err error) {
	if err = r.Db.Model(&domain.Member{}).Joins(`
																					LEFT JOIN member_travel ON member_travel.member_id = members.id
																					LEFT JOIN travels ON member_travel.travel_id = travels.id
																	`).Where("travels.travel_key = ?", travel_key).Find(&members).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find member  - %s", err)
		return
	}
	return
}

func (r *MemberRepository) FindByTravelKeyWithBorrowMoneyListAndPayments(travel_key string) (members domain.Members, err error) {
	if err = r.Db.Preload("BorrowMoneyList").Preload("Payments").Joins(`
																					LEFT JOIN member_travel ON member_travel.member_id = members.id
																					LEFT JOIN travels ON member_travel.travel_id = travels.id
																	`).Where("travels.travel_key = ?", travel_key).Find(&members).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find member  - %s", err)
		return
	}
	return
}

func (r *MemberRepository) FindByIdWithBorrowMoneyListAndPayments(member_id int) (member domain.Member, err error) {
	if err = r.Db.Preload("BorrowMoneyList").Preload("Payments").Where(domain.Member{ID: uint(member_id)}).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = domain.Errorf(codes.NotFound, "Failed to find member - %s", err)
			return
		}
		err = domain.Errorf(codes.Database, "Failed to find member  - %s", err)
		return
	}
	return
}

func (r *MemberRepository) FindById(id int) (member domain.Member, err error) {
	if err = r.Db.First(&member, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = domain.Errorf(codes.NotFound, "Failed to find member - %s", err)
			return
		}
		err = domain.Errorf(codes.Database, "Failed to find member  - %s", err)
		return
	}
	return
}

func (r *MemberRepository) FindByPaymentId(payment_id int) (members domain.Members, err error) {
	if err = r.Db.Model(&domain.Member{}).Joins(`
																					LEFT JOIN borrow_money ON borrow_money.borrower_id = members.id
																	`).Where("borrow_money.payment_id = ?", payment_id).Find(&members).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find member  - %s", err)
		return
	}
	return
}

func (r *MemberRepository) StoreMembers(ctx context.Context, members domain.Members) (created_members domain.Members, err error) {
	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&members).Error
	} else {
		err = r.Db.Create(&members).Error
	}

	created_members = members

	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create member  - %s", err)
		return
	}
	return
}

func (r *MemberRepository) StoreMember(ctx context.Context, member domain.Member) (created_member domain.Member, err error) {
	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&member).Scan(&created_member).Error
	} else {
		err = r.Db.Create(&member).Scan(&created_member).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create member  - %s", err)
		return
	}
	return
}

func (r *MemberRepository) DeleteMembers(ctx context.Context, members domain.Members) (err error) {
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

	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&deleted_members).Error
	} else {
		err = r.Db.Create(&deleted_members).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_member  - %s", err)
		return
	}

	member_ids := []int{}
	for _, member := range members {
		member_ids = append(member_ids, int(member.ID))
	}

	if ok {
		err = db.Where("id IN (?)", member_ids).Delete(&domain.Member{}).Error
	} else {
		err = r.Db.Where("id IN (?)", member_ids).Delete(&domain.Member{}).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete member  - %s", err)
		return
	}
	return
}

func (r *MemberRepository) DeleteMember(ctx context.Context, member domain.Member) (err error) {
	deleted_member := domain.DeletedMember{
		ID:        member.ID,
		Name:      member.Name,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
		DeletedAt: time.Now(),
	}

	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&deleted_member).Error
	} else {
		err = r.Db.Create(&deleted_member).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_member  - %s", err)
		return
	}

	if ok {
		err = db.Model(&member).Delete(&member).Error
	} else {
		err = r.Db.Model(&member).Delete(&member).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete member  - %s", err)
		return
	}
	return
}
