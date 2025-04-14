package merchant

import (
	"errors"
	"fmt"
	"strings"
	"wkm/utils"

	"gorm.io/gorm"
)

type MerchantRepository interface {
	CreateMerchant(data Merchant) error
	MasterData(search string, kategori string, lokasi string, limit int, pageParams int) []Merchant
	MasterDataSearch(search string) []Merchant
	MasterDataCount(search string, kategori string, lokasi string) int64
	MasterDataAll() []Merchant
	DetailMerchant(id string, lokasi string) Merchant
	Delete(id string, name string) error
	Update(body Merchant) error
}

type merchantRepository struct {
	conn *gorm.DB
}

func NewMerchantRepository(conn *gorm.DB) MerchantRepository {
	return &merchantRepository{
		conn: conn,
	}
}

func (lR *merchantRepository) DetailMerchant(id string, lokasi string) Merchant {
	merchant := Merchant{ID: id}
	query := lR.conn.Preload("Kategori").Preload("MediaPromosi").Preload("NamaPICMRO")
	if lokasi != "" {
		query.Preload("Outlet", "kota in ?", strings.Split(lokasi, "w3"))
	} else {
		query.Preload("Outlet")
	}
	query.Find(&merchant)
	return merchant
}

func (lR *merchantRepository) CreateMerchant(data Merchant) error {
	var count int64
	lR.conn.Model(&Merchant{}).Where("pin = 1").Count(&count)
	if count >= 18 && data.Pin {
		return errors.New("maaf merchant sudah mencapai batas maksimal")
		
	}
	kategori := data.Kategori
	mediaPromosi := data.MediaPromosi
	picMro := data.NamaPICMRO
	result := lR.conn.Omit("Kategori", "MediaPromosi", "NamaPICMRO").Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	err := lR.conn.Model(&data).Association("Kategori").Replace(kategori)
	if err != nil {
		return err
	}
	err = lR.conn.Model(&data).Association("MediaPromosi").Replace(mediaPromosi)
	if err != nil {
		return err
	}
	err = lR.conn.Model(&data).Association("NamaPICMRO").Replace(picMro)
	if err != nil {
		return err
	}
	return nil

}

func (lR *merchantRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&Merchant{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
        "updated_by": name, // kalau kamu mau set waktu update juga
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (lR *merchantRepository) Update(data Merchant) error {
	record := Merchant{ID: data.ID}
	var count int64
	lR.conn.Model(&Merchant{}).Where("pin = 1").Count(&count)
	lR.conn.First(&record)
	if record.Nama == "" {
		return errors.New("maaf data tidak ada")
	}
	if count >= 18 && !record.Pin && data.Pin {
		return errors.New("maaf merchant sudah mencapai batas maksimal")
		
	}
	kategori := data.Kategori
	mediaPromosi := data.MediaPromosi
	picMro := data.NamaPICMRO

	result := lR.conn.Omit("Kategori", "MediaPromosi", "NamaPICMRO").Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	err := lR.conn.Model(&record).Association("Kategori").Replace(kategori)
	if err != nil {
		return err
	}
	err = lR.conn.Model(&record).Association("MediaPromosi").Replace(mediaPromosi)
	if err != nil {
		return err
	}
	err = lR.conn.Model(&record).Association("NamaPICMRO").Replace(picMro)
	if err != nil {
		return err
	}
	return nil
}

func (lR *merchantRepository) MasterData(search string, kategori string, lokasi string, limit int, pageParams int) []Merchant {
	merchant := []Merchant{}
	query := lR.conn.Select("DISTINCT merchants.id, merchants.nama, merchants.logo, merchants.is_active, merchants.nama_pic, merchants.no_telp_pic, merchants.website, merchants.email, merchants.pin,merchants.updated_at").Where("merchants.nama like ? or merchants.alamat like  ? ", "%"+search+"%", "%"+search+"%").Where("is_deleted = 0")
	if kategori != "" {
		query.Joins("JOIN merchant_kategoris a ON a.MerchantID = merchants.id").
			Joins("JOIN mst_kategori b ON a.KategoriID = b.id").
			Where("b.nama in ?", strings.Split(kategori, "w3"))
	}
	if lokasi != "" {
		query.Joins("JOIN outlets c ON c.merchant_id = merchants.id").
			Where("c.kota in ?", strings.Split(lokasi, "w3"))
	}
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Order("merchants.pin desc,merchants.updated_at desc").Find(&merchant)
	return merchant
}

func (lR *merchantRepository) MasterDataSearch(search string) []Merchant {
	merchant := []Merchant{}
	query := lR.conn.Select("DISTINCT merchants.id, merchants.nama").Where("merchants.is_deleted = 0")
	if search != "" {
		query.Joins("JOIN merchant_kategoris a ON a.MerchantID = merchants.id").
			Joins("JOIN mst_kategori b ON a.KategoriID = b.id").
			Joins("JOIN outlets c ON c.merchant_id = merchants.id").
			Where("merchants.nama like ? or merchants.alamat like ? or b.nama like ? or c.kota like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	query.Limit(3).Find(&merchant)
	return merchant
}

func (lR *merchantRepository) MasterDataCount(search string, kategori string, lokasi string) int64 {
	var merchant []Merchant
	query := lR.conn.Where("nama like ? or alamat like  ? ", "%"+search+"%", "%"+search+"%").Where("is_deleted = 0")
	if kategori != "" {
		query.Joins("JOIN merchant_kategoris a ON a.MerchantID = merchants.id").
			Joins("JOIN mst_kategori b ON a.KategoriID = b.id").
			Where("b.nama in ?", strings.Split(kategori, "w3"))
	}
	if lokasi != "" {
		query.Joins("JOIN outlets c ON c.merchant_id = merchants.id").
			Where("c.kota in ?", strings.Split(lokasi, "w3"))
	}
	query.Select("DISTINCT merchants.id").Find(&merchant)
	return int64(len(merchant))
}

func (lR *merchantRepository) MasterDataAll() []Merchant {
	var merchant []Merchant
	lR.conn.Select("id, nama").Where("is_deleted = 0").Find(&merchant)
	return merchant
}
