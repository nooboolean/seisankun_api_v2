package usecases

import (
	"math/rand"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
)

type TravelInteractor struct {
	TravelRepository       *repositories.TravelRepository
	MemberRepository       *repositories.MemberRepository
	MemberTravelRepository *repositories.MemberTravelRepository
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

	// NOTE: member_travelのmember_idやtravel_idに外部キー制約がかかっているので、先に削除の必要あり
	err = i.MemberTravelRepository.DeleteList(member_travel_list)
	if err != nil {
		return
	}
	err = i.TravelRepository.Delete(travel)
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
