package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string   `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	NoHp        string   `form:"no_hp" json:"no_hp" gorm:"column:no_hp"`
	Name        string   `form:"name" json:"name" gorm:"column:name"`
	Password    string   `form:"password" json:"password" gorm:"column:password"`
	Permissions []string `gorm:"type:text;->"`
	Active      bool     `json:"active" gorm:"column:active"`
}

func (User) TableName() string {
	return "users"
}

func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

type Role struct {
	ID   string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Name string `form:"name" json:"name" gorm:"column:name"`
}

func (Role) TableName() string {
	return "mst_roles"
}

type Permission struct {
	ID     string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Name   string `form:"name" json:"name" gorm:"column:name"`
	RoleId string `form:"role_id" json:"role_id" gorm:"column:role_id"`
}

func (Permission) TableName() string {
	return "mst_permissions"
}
