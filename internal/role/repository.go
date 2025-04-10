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
	DetailRole(id uint64) entity.Role
	Update(body entity.Role) error
	Delete(id string, name string) error
}

type roleRepository struct {
	conn *gorm.DB
}

func NewRoleRepository(conn *gorm.DB) RoleRepository {
	return &roleRepository{
		conn: conn,
	}
}

func (lR *roleRepository) MasterDataAll() []entity.Role {
	var roles []entity.Role
	lR.conn.Select("id, name").Where("is_deleted = 0").Find(&roles)
	return roles
}

func (lR *roleRepository) DetailRole(id uint64) entity.Role {
	user := entity.Role{ID: id}
	lR.conn.Preload("Permissions").Find(&user)
	return user
}

func (lR *roleRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&entity.Role{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *roleRepository) CreateRole(data entity.Role) error {
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *roleRepository) Update(data entity.Role) error {
	var record entity.Role
	if err := lR.conn.First(&record, data.ID).Error; err != nil {
		return errors.New("maaf data tidak ada")
	}

	// Update nama role, tanpa update Permissions
	if err := lR.conn.Model(&record).Update("name", data.Name).Error; err != nil {
		return err
	}

	// Hapus semua permissions lama
	if err := lR.conn.Where("role_id = ?", data.ID).Delete(&entity.Permission{}).Error; err != nil {
		return err
	}

	// Buat permission baru (name dari client, role_id dari data.ID)
	var newPermissions []entity.Permission
	for _, p := range data.Permissions {
		newPermissions = append(newPermissions, entity.Permission{
			Name:   p.Name,
			RoleId: data.ID, // karena RoleId kamu tipe string
		})
	}

	if len(newPermissions) > 0 {
		if err := lR.conn.Create(&newPermissions).Error; err != nil {
			return err
		}
	}

	return nil
}

func (lR *roleRepository) MasterData(search string, limit int, pageParams int) []entity.Role {
	user := []entity.Role{}
	query := lR.conn.Where("is_deleted = 0").Where("name like ? ", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&user)
	return user
}

func (lR *roleRepository) MasterDataCount(search string) int64 {
	var user []entity.Role
	query := lR.conn.Where("is_deleted = 0").Where("name like ? ", "%"+search+"%")
	query.Select("id").Find(&user)
	return int64(len(user))
}
