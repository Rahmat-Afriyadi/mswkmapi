package user

import "wkm/entity"

type UserSService interface {
	CreateUserS(data entity.UserS) error
	MasterData(search string, limit int, pageParams int) []entity.UserS
	MasterDataCount(search string) int64
	DetailUserS(id uint64) entity.UserS
	Update(body entity.UserS) error
	Delete(id string, name string) error
}

type userSService struct {
	trR UserSRepository
}

func NewUserSService(tR UserSRepository) UserSService {
	return &userSService{
		trR: tR,
	}
}

func (s *userSService) Update(body entity.UserS) error {
	return s.trR.Update(body)
}
func (s *userSService) Delete(id string, name string) error {
	return s.trR.Delete(id, name)
}
func (s *userSService) DetailUserS(id uint64) entity.UserS {
	return s.trR.DetailUserS(id)
}
func (s *userSService) MasterData(search string, limit int, pageParams int) []entity.UserS {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *userSService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *userSService) CreateUserS(data entity.UserS) error {
	return s.trR.CreateUserS(data)
}
