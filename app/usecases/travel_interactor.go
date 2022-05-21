package usecases

import (
	"math/rand"
	"time"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
)

type TravelInteractor struct {
	TravelRepository              *repositories.TravelRepository
	DeletedTravelRepository       *repositories.DeletedTravelRepository
	MemberRepository              *repositories.MemberRepository
	DeletedMemberRepository       *repositories.DeletedMemberRepository
	MemberTravelRepository        *repositories.MemberTravelRepository
	DeletedMemberTravelRepository *repositories.DeletedMemberTravelRepository
}

func (i *TravelInteractor) Get(travel_key string) (travel domain.Travel, members domain.Members, err error) {
	travel, err = i.TravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return travel, domain.Members{}, err
	}
	members, err = i.MemberRepository.FindByTravelKey(travel_key)
	if err != nil {
		return domain.Travel{}, members, err
	}
	return
}

func (i *TravelInteractor) Register(members domain.Members, travel domain.Travel) (travel_key string, err error) {
	travel.TravelKey = RandomString(30)
	travel_key, err = i.TravelRepository.Store(&travel)
	if err != nil {
		return travel_key, err
	}
	err = i.MemberRepository.StoreMembers(members)
	if err != nil {
		return travel_key, err
	}

	member_travel_list := make(domain.MemberTravelList, 0, len(members))
	for _, member := range members {
		member_travel := domain.MemberTravel{
			MemberId: int(member.ID),
			TravelId: int(travel.ID),
		}
		member_travel_list = append(member_travel_list, member_travel)
	}

	err = i.MemberTravelRepository.StoreList(member_travel_list)
	return
}

func (i *TravelInteractor) Update(t domain.Travel) (travel domain.Travel, err error) {
	travel, err = i.TravelRepository.Update(t)
	if err != nil {
		return travel, err
	}
	return
}

func (i *TravelInteractor) Delete(travel_key string) (err error) {
	travel, err := i.TravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}

	members, err := i.MemberRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}

	member_travel_list, err := i.MemberTravelRepository.FindByTravelKey(travel_key)
	if err != nil {
		return
	}

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
	err = i.DeletedMemberTravelRepository.StoreList(deleted_member_travel_list)
	if err != nil {
		return
	}

	err = i.MemberTravelRepository.DeleteList(member_travel_list)
	if err != nil {
		return
	}

	err = i.DeletedTravelRepository.Store(travel)
	if err != nil {
		return
	}
	err = i.TravelRepository.Delete(travel.TravelKey)
	if err != nil {
		return
	}
	err = i.DeletedMemberRepository.StoreMembers(members)
	if err != nil {
		return
	}
	err = i.MemberRepository.DeleteMembers(members)
	if err != nil {
		return
	}
	return
}

//TODO: ユニークではないため、なにかのモジュールを使って対応する
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
