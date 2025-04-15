package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kategori struct {
	ID   string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Nama string `form:"nama" json:"name" gorm:"type:varchar(100);column:nama"`
	IsDeleted     bool                  `form:"is_deleted" json:"is_deleted" gorm:"column:is_deleted;default:false"`

}

func (Kategori) TableName() string {
	return "mst_kategori"
}

func (b *Kategori) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
