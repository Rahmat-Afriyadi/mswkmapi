package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Outlet struct {
	ID           string         `form:"id" json:"id" gorm:"type:varchar(36);primary_key;column:id"`
	Nama         string         `form:"nama" json:"nama" gorm:"type:varchar(100);column:nama"`
	Alamat       string         `form:"alamat" json:"alamat" gorm:"column:alamat"`
	Latitude     string         `form:"latitude" json:"latitude" gorm:"column:latitude;type:decimal(3,15)"`
	Longitude    string         `form:"longitude" json:"longitude" gorm:"column:longitude;type:decimal(3,15)"`
	NamaPIC      string         `form:"nama_pic" json:"nama_pic" gorm:"type:varchar(100);column:nama_pic"`
	NoTelpPIC    string         `form:"no_telp_pic" json:"no_telp_pic" gorm:"type:varchar(15);column:no_telp_pic"`
	IsActive     bool           `form:"is_active" json:"is_active" gorm:"column:is_active;default:true"`
	MerchantId   string         `form:"merchant_id" json:"merchant_id" gorm:"column:merchant_id;"`
	Merchant     Merchant       `form:"merchant" json:"merchant" gorm:"->;references:ID;foreignKey:MerchantId"`
	MediaPromosi []MediaPromosi `json:"media_promosi" gorm:"many2many:outlet_media_promosi;"`
	CreatedBy    string         `form:"created_by" json:"created_by" gorm:"column:created_by"`
	UpdatedBy    string         `form:"updated_by" json:"updated_by" gorm:"column:updated_by"`
	CreatedAt    *time.Time     `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time     `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Outlet) TableName() string {
	return "outlets"
}

func (b *Outlet) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
