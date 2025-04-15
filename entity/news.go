package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	ID           string         `form:"id" json:"id" gorm:"type:varchar(36);primary_key;column:id"`
	Nama         string         `form:"nama" json:"nama" gorm:"type:varchar(100);column:nama"`
	IsActive     bool           `form:"is_active" json:"is_active" gorm:"column:is_active;default:true"`
	Pin     bool           `form:"pin" json:"pin" gorm:"column:pin;default:false"`
	IsDeleted     bool                  `form:"is_deleted" json:"is_deleted" gorm:"column:is_deleted;default:false"`
	Kategori     []NewsKategori     `json:"kategori" gorm:"many2many:news_kategori;association_autocreate:false;"`
	Deskripsi    string         `form:"deskripsi" json:"deskripsi" gorm:"column:deskripsi;"`
	Logo         string         `form:"logo" json:"logo" gorm:"type:varchar(100);column:logo"`
	Banner       string         `form:"banner" json:"banner" gorm:"type:varchar(100);column:banner"`
	CreatedBy    string         `form:"created_by" json:"created_by" gorm:"type:varchar(100);column:created_by"`
	UpdatedBy    string         `form:"updated_by" json:"updated_by" gorm:"type:varchar(100);column:updated_by"`
	CreatedAt    *time.Time     `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time     `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

}

func (News) TableName() string {
	return "news"
}

func (b *News) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
