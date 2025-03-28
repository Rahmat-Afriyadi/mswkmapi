package repository

import (
	"errors"
	"wkm/entity"
	"wkm/request"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id string) entity.User
	FindByPhoneNumber(username string) entity.User
	ResetPassword(data request.ResetPassword)
	CreateUser(data request.SignupRequest) (entity.User, error)
	FindByUsername(username string) entity.UserS
	FindByIdAdmin(id string) entity.UserS
}

type userRepository struct {
	connUser *gorm.DB
}

func NewUserRepository(connUser *gorm.DB) UserRepository {
	return &userRepository{
		connUser: connUser,
	}
}

func (lR *userRepository) FindById(id string) entity.User {
	user := entity.User{ID: id}
	lR.connUser.Find(&user)

	return user
}

func (lR *userRepository) ResetPassword(data request.ResetPassword) {
	user := entity.User{ID: data.IdUser}
	lR.connUser.Find(&user)
	user.Password = data.Password
	lR.connUser.Save(&user)
}

func (lR *userRepository) CreateUser(data request.SignupRequest) (entity.User, error) {
	user := entity.User{NoHp: data.NoHp}
	lR.connUser.Where("no_hp", user.NoHp).First(&user)
	if user.ID != "" {
		return user, errors.New("nomor tersebut telah terdaftar")
	}
	user.Name = data.Fullname
	user.Password = data.Password
	result := lR.connUser.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (lR *userRepository) FindByUsername(username string) entity.UserS {
	var user entity.UserS
	lR.connUser.Where("username", username).First(&user)

	var permissions []entity.Permission
	lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	for _, v := range permissions {
		user.Permissions = append(user.Permissions, v.Name)
	}

	return user
}

func (lR *userRepository) FindByIdAdmin(id string) entity.UserS {
	var user entity.UserS
	lR.connUser.Where("id", id).First(&user)

	var permissions []entity.Permission
	lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	for _, v := range permissions {
		user.Permissions = append(user.Permissions, v.Name)
	}

	return user
}

func (lR *userRepository) FindByPhoneNumber(username string) entity.User {
	var user entity.User
	lR.connUser.Where("no_hp", username).First(&user)

	return user
}
