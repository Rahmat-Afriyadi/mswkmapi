package user

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type UserSRepository interface {
	CreateUserS(data entity.UserS) error
	MasterData(search string, limit int, pageParams int) []entity.UserS
	MasterDataCount(search string) int64
	DetailUserS(id string) entity.UserS
	Update(body entity.UserS) error
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserSRepository(conn *gorm.DB) UserSRepository {
	return &userRepository{
		conn: conn,
	}
}

func (lR *userRepository) DetailUserS(id string) entity.UserS {
	user := entity.UserS{ID: id}
	lR.conn.Find(&user)
	fmt.Println("ini user yaa ", user)
	return user
}

func (lR *userRepository) CreateUserS(data entity.UserS) error {
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
	query := lR.conn.Where("name like ? ", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&user)
	return user
}

func (lR *userRepository) MasterDataCount(search string) int64 {
	var user []entity.UserS
	query := lR.conn.Where("name like ? ", "%"+search+"%")
	query.Select("id").Find(&user)
	return int64(len(user))
}
