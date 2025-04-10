package user

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSRepository interface {
	CreateUserS(data entity.UserS) error
	MasterData(search string, limit int, pageParams int) []entity.UserS
	MasterDataCount(search string) int64
	DetailUserS(id uint64) entity.UserS
	Update(body entity.UserS) error
	Delete(id string, name string) error
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserSRepository(conn *gorm.DB) UserSRepository {
	return &userRepository{
		conn: conn,
	}
}

func (lR *userRepository) DetailUserS(id uint64) entity.UserS {
	user := entity.UserS{ID: id}
	lR.conn.Find(&user)
	return user
}
func (lR *userRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&entity.UserS{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *userRepository) CreateUserS(data entity.UserS) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	data.Password = string(password)
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *userRepository) Update(data entity.UserS) error {
	record := entity.UserS{ID: data.ID}
	lR.conn.Find(&record)
	if record.Name == "" {
		return errors.New("maaf data tidak ada")
	}
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *userRepository) MasterData(search string, limit int, pageParams int) []entity.UserS {
	user := []entity.UserS{}
	query := lR.conn.Preload("Role").Where("name like ? ", "%"+search+"%").Where("is_deleted = 0")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&user)
	fmt.Println(user[0].ID, user[0].Role.Name)
	return user
}

func (lR *userRepository) MasterDataCount(search string) int64 {
	var user []entity.UserS
	query := lR.conn.Where("name like ? ", "%"+search+"%").Where("is_deleted = 0")
	query.Select("id").Find(&user)
	return int64(len(user))
}
