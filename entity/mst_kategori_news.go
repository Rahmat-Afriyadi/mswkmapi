package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KategoriNews struct {
	ID   string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Nama string `form:"nama" json:"name" gorm:"type:varchar(100);column:nama"`
	IsDeleted     bool                  `form:"is_deleted" json:"is_deleted" gorm:"column:is_deleted;default:false"`

}

func (KategoriNews) TableName() string {
	return "mst_kategori_news"
}

func (b *KategoriNews) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
