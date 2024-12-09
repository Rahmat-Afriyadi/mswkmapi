package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kategori struct {
	ID   string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Nama string `form:"nama" json:"nama" gorm:"type:varchar(100);column:nama"`
}

func (Kategori) TableName() string {
	return "mst_kategori"
}

func (b *Kategori) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
