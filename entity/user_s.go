package entity

import "time"

type UserS struct {
	ID          uint64     `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	Username    string     `form:"username" json:"username" gorm:"column:username"`
	Name        string     `form:"name" json:"name" gorm:"column:name"`
	Password    string     `form:"password" json:"password" gorm:"column:password2"`
	Permissions []string   `gorm:"type:text;->"`
	RoleId      uint32     `form:"role_id" json:"role_id" gorm:"column:role_id"`
	Role        Role       `form:"role" json:"role" gorm:"->;references:ID;foreignKey:RoleId"`
	CreatedBy   *string     `form:"created_by" json:"created_by" gorm:"column:created_by"`
	UpdatedBy   *string     `form:"updated_by" json:"updated_by" gorm:"column:updated_by"`
	CreatedAt   *time.Time `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (UserS) TableName() string {
	return "userses"
}

type Role struct {
	ID   uint64 `form:"id" json:"id" gorm:"primary_key;column:id"`
	Name string `form:"name" json:"name" gorm:"column:name"`
	Permissions []Permission `json:"permissions" gorm:"foreignKey:RoleId;references:ID"`
}

func (Role) TableName() string {
	return "mst_roles"
}

type Permission struct {
	ID     uint64 `form:"id" json:"id" gorm:"primary_key;column:id"`
	Name   string `form:"name" json:"name" gorm:"column:name"`
	RoleId uint64 `form:"role_id" json:"role_id" gorm:"column:role_id"`
}

func (Permission) TableName() string {
	return "mst_permissions"
}