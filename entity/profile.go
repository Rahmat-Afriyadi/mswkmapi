package entity

import "time"

type Profile struct {
	Id         string     `form:"id" json:"id" gorm:"primary_key;column:id"`
	NoHp       string     `form:"no_hp" json:"no_hp" gorm:"column:no_hp"`
	Name       string     `form:"name" json:"name" gorm:"column:name"`
	Email      string     `form:"email" json:"email" gorm:"column:email"`
	ImgProfile string     `form:"img_profile" json:"img_profile" gorm:"column:img_profile"`
	JnsKelamin string     `form:"jns_kelamin" json:"jns_kelamin" gorm:"column:jns_kelamin"`
	TglLahir   *time.Time `form:"tgl_lahir" json:"tgl_lahir" gorm:"column:tgl_lahir"`
	Alamat     string     `form:"alamat" json:"alamat" gorm:"column:alamat"`
}

func (Profile) TableName() string {
	return "users"
}
