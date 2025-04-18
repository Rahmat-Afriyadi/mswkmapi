package entity

import (
	"time"
)

type Member struct {
	NoMsn      string    `form:"no_msn" json:"no_msn" gorm:"column:no_msn"`
	NmCustomer string    `form:"nm_customer" json:"nm_customer" gorm:"column:nm_customer11"`
	NoKartu    string    `form:"no_kartu" json:"no_kartu" gorm:"primary_key;column:no_kartu"`
	TglExpired time.Time `form:"tgl_expired" json:"tgl_expired" gorm:"column:tgl_expired"`
	UserId     *string   `json:"user_id" gorm:"column:user_id;"`
	User       User      `json:"user" gorm:"references:ID;foreignKey:NoMsn"`
}

func (Member) TableName() string {
	return "member"
}
