package entity

type Kodepos struct {
	KdPos     string `json:"kd_pos" gorm:"column:kd_pos"`
	Kodepos   string `json:"kodepos" gorm:"column:kodepos"`
	Kelurahan string `json:"kelurahan" gorm:"column:kelurahan"`
	Kecamatan string `json:"kecamatan" gorm:"column:kecamatan"`
	Kota      string `json:"kota" gorm:"column:kota"`
}

func (Kodepos) TableName() string {
	return "kodepos"
}
