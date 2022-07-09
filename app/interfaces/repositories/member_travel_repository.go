package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"gorm.io/gorm"
)

type MemberTravelRepository struct {
	Db SqlHandler
}

func (r *MemberTravelRepository) FindByTravelKey(travel_key string) (memberTravelList domain.MemberTravelList, err error) {
	if err = r.Db.Model(&domain.MemberTravel{}).Joins(`
																								LEFT JOIN travels ON member_travel.travel_id = travels.id
																				`).Where("travels.travel_key = ?", travel_key).Find(&memberTravelList).Error; err != nil {
		err = domain.Errorf(codes.Database, "Failed to find member_travel  - %s", err)
		return
	}
	return
}

func (r *MemberTravelRepository) FindByMemberId(member_id int) (memberTravel domain.MemberTravel, err error) {
	if err = r.Db.Model(&domain.MemberTravel{}).Where("member_id = ?", member_id).First(&memberTravel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = domain.Errorf(codes.NotFound, "Failed to find member_travel - %s", err)
			return
		}
		err = domain.Errorf(codes.Database, "Failed to find member_travel  - %s", err)
		return
	}
	return
}

func (r *MemberTravelRepository) FindByMemberIdAndTravelId(member_id int, travel_id int) (memberTravel domain.MemberTravel, err error) {
	if err = r.Db.Model(&domain.MemberTravel{}).Where("member_id = ?", member_id).Where("travel_id = ?", travel_id).First(&memberTravel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = domain.Errorf(codes.NotFound, "Failed to find member_travel - %s", err)
			return
		}
		err = domain.Errorf(codes.Database, "Failed to find member_travel  - %s", err)
		return
	}
	return
}

func (r *MemberTravelRepository) StoreList(ctx context.Context, memberTravelList domain.MemberTravelList) (err error) {
	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&memberTravelList).Error
	} else {
		err = r.Db.Create(&memberTravelList).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create member_travel  - %s", err)
		return
	}
	return
}

func (r *MemberTravelRepository) Store(ctx context.Context, memberTravel domain.MemberTravel) (err error) {
	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&memberTravel).Error
	} else {
		err = r.Db.Create(&memberTravel).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create member_travel  - %s", err)
		return
	}
	return
}

func (r *MemberTravelRepository) DeleteList(ctx context.Context, member_travel_list domain.MemberTravelList) (err error) {
	deleted_member_travel_list := make(domain.DeletedMemberTravelList, 0, len(member_travel_list))
	deleted_at := time.Now()
	for _, member_travel := range member_travel_list {
		deleted_member_travel := domain.DeletedMemberTravel{
			ID:        member_travel.ID,
			MemberId:  member_travel.MemberId,
			TravelId:  member_travel.TravelId,
			CreatedAt: member_travel.CreatedAt,
			DeletedAt: deleted_at,
		}
		deleted_member_travel_list = append(deleted_member_travel_list, deleted_member_travel)
	}

	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&deleted_member_travel_list).Error
	} else {
		err = r.Db.Create(&deleted_member_travel_list).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_member_travel  - %s", err)
		return
	}

	member_travel_list_ids := []int{}
	for _, member_travel := range member_travel_list {
		member_travel_list_ids = append(member_travel_list_ids, int(member_travel.ID))
	}

	if ok {
		err = db.Where("id IN (?)", member_travel_list_ids).Delete(&domain.MemberTravel{}).Error
	} else {
		err = r.Db.Where("id IN (?)", member_travel_list_ids).Delete(&domain.MemberTravel{}).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete member_travel  - %s", err)
		return
	}
	return
}

func (r *MemberTravelRepository) Delete(ctx context.Context, member_travel domain.MemberTravel) (err error) {
	deleted_member_travel := domain.DeletedMemberTravel{
		ID:        member_travel.ID,
		MemberId:  member_travel.MemberId,
		TravelId:  member_travel.TravelId,
		CreatedAt: member_travel.CreatedAt,
		DeletedAt: time.Now(),
	}

	db, ok := GetTx(ctx)
	if ok {
		err = db.Create(&deleted_member_travel).Error
	} else {
		err = r.Db.Create(&deleted_member_travel).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to create deleted_member_travel  - %s", err)
		return
	}

	if ok {
		err = db.Model(&member_travel).Delete(&member_travel).Error
	} else {
		err = r.Db.Model(&member_travel).Delete(&member_travel).Error
	}
	if err != nil {
		err = domain.Errorf(codes.Database, "Failed to delete member_travel  - %s", err)
		return
	}
	return
}
