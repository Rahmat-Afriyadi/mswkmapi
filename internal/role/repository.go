package role

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(data entity.Role) error
	MasterDataAll() []entity.Role
	MasterData(search string, limit int, pageParams int) []entity.Role
	MasterDataCount(search string) int64
	DetailRole(id string) entity.Role
	Update(body entity.Role) error
}

type userRepository struct {
	conn *gorm.DB
}

func NewRoleRepository(conn *gorm.DB) RoleRepository {
	return &userRepository{
		conn: conn,
	}
}

func (lR *userRepository) MasterDataAll() []entity.Role {
	var roles []entity.Role
	lR.conn.Select("id, name").Find(&roles)
	return roles
}

func (lR *userRepository) DetailRole(id string) entity.Role {
	user := entity.Role{ID: id}
	lR.conn.Find(&user)
	fmt.Println("ini user yaa ", user)
	return user
}

func (lR *userRepository) CreateRole(data entity.Role) error {
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *userRepository) Update(data entity.Role) error {
	record := entity.Role{ID: data.ID}
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

func (lR *userRepository) MasterData(search string, limit int, pageParams int) []entity.Role {
	user := []entity.Role{}
	query := lR.conn.Where("name like ? ", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&user)
	return user
}

func (lR *userRepository) MasterDataCount(search string) int64 {
	var user []entity.Role
	query := lR.conn.Where("name like ? ", "%"+search+"%")
	query.Select("id").Find(&user)
	return int64(len(user))
}
