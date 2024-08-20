package repository

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/request"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id string) entity.User
	FindByPhoneNumber(username string) entity.User
	ResetPassword(data request.ResetPassword)
	CreateUser(data request.SignupRequest) (entity.User, error)
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

	// var permissions []entity.Permission
	// lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	// for _, v := range permissions {
	// 	user.Permissions = append(user.Permissions, v.Name)
	// }

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
	fmt.Println("ini user yaa ", user)
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

func (lR *userRepository) FindByPhoneNumber(username string) entity.User {
	var user entity.User
	lR.connUser.Where("no_hp", username).First(&user)

	// var permissions []entity.Permission
	// lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	// for _, v := range permissions {
	// 	user.Permissions = append(user.Permissions, v.Name)
	// }

	return user
}
