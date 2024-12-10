package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaPromosi struct {
	ID   string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Nama string `form:"nama" json:"name" gorm:"type:varchar(100);column:nama"`
}

func (MediaPromosi) TableName() string {
	return "mst_media_promosi"
}

func (b *MediaPromosi) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
