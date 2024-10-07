package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	Me(userId string) entity.Profile
	Update(user entity.Profile) entity.Profile
}

type profileRepository struct {
	connUser *gorm.DB
}

func NewProfileRepository(connUser *gorm.DB) ProfileRepository {
	return &profileRepository{
		connUser: connUser,
	}
}

func (s *profileRepository) Me(userId string) entity.Profile {
	data := entity.Profile{}
	s.connUser.Where("id", userId).Find(&data)
	return data
}
func (s *profileRepository) Update(user entity.Profile) entity.Profile {
	s.connUser.Save(&user)
	return user
}
