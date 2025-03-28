package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string   `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	NoHp        string   `form:"no_hp" json:"no_hp" gorm:"column:no_hp"`
	IsAdmin     bool     `form:"is_admin" json:"is_admin" gorm:"column:is_admin"`
	Email       string   `form:"email" json:"email" gorm:"column:email"`
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

