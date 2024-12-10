package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PicMro struct {
	ID   string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Nama string `form:"nama" json:"name" gorm:"type:varchar(100);column:nama"`
}

func (PicMro) TableName() string {
	return "mst_pic_mro"
}

func (b *PicMro) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
