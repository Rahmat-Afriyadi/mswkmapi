package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Otp struct {
	ID        string     `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	NoHp      string     `form:"no_hp" json:"no_hp" gorm:"column:no_hp"`
	Used      bool       `form:"used" json:"used" gorm:"column:used"`
	Otp       int        `form:"otp" json:"otp" gorm:"column:otp"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

func (Otp) TableName() string {
	return "otp"
}

func (b *Otp) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
