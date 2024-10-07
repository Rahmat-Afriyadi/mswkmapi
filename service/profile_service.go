package service

import (
	"wkm/entity"
	"wkm/repository"
)

type ProfileService interface {
	Me(userId string) entity.Profile
	Update(user entity.Profile) entity.Profile
}

type profileService struct {
	pR repository.ProfileRepository
}

func NewProfileService(pR repository.ProfileRepository) ProfileService {
	return &profileService{
		pR,
	}
}

func (s *profileService) Me(userId string) entity.Profile {
	return s.pR.Me(userId)
}
func (s *profileService) Update(user entity.Profile) entity.Profile {
	return s.pR.Update(user)
}
