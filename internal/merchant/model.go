package merchant

import (
	"time"
	"wkm/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Merchant struct {
	ID           string                `form:"id" json:"id" gorm:"type:varchar(36);primary_key;column:id"`
	Nama         string                `form:"nama" json:"nama" gorm:"type:varchar(100);column:nama"`
	Alamat       string                `form:"alamat" json:"alamat" gorm:"column:alamat"`
	NamaPIC      string                `form:"nama_pic" json:"nama_pic" gorm:"type:varchar(100);column:nama_pic"`
	NoTelpPIC    string                `form:"no_telp_pic" json:"no_telp_pic" gorm:"type:varchar(15);column:no_telp_pic"`
	IsActive     bool                  `form:"is_active" json:"is_active" gorm:"column:is_active;default:true"`
	ValidFrom    time.Time             `form:"valid_from" json:"valid_from" gorm:"type:DATE;column:valid_from"`
	ValidThru    time.Time             `form:"valid_thru" json:"valid_thru" gorm:"type:DATE;column:valid_thru"`
	Logo         string                `form:"logo" json:"logo" gorm:"type:varchar(100);column:logo"`
	Kategori     []entity.Kategori     `json:"kategori" gorm:"many2many:merchant_kategori;association_autocreate:false;"`
	Email        string                `form:"email" json:"email" gorm:"type:varchar(100);column:email"`
	Website      string                `form:"website" json:"website" gorm:"type:varchar(100);column:website"`
	MediaPromosi []entity.MediaPromosi `json:"media_promosi" gorm:"many2many:merchant_media_promosi;association_autocreate:false;"`
	Deskripsi    string                `form:"deskripsi" json:"deskripsi" gorm:"column:deskripsi;"`
	NamaPICMRO   []entity.PicMro       `json:"nama_pic_mro" gorm:"many2many:merchant_pic_mro;association_autocreate:false;"`
	CreatedBy    string                `form:"created_by" json:"created_by" gorm:"type:varchar(100);column:created_by"`
	UpdatedBy    string                `form:"updated_by" json:"updated_by" gorm:"type:varchar(100);column:updated_by"`
	CreatedAt    *time.Time            `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time            `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Merchant) TableName() string {
	return "merchants"
}

func (b *Merchant) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
