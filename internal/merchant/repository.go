package merchant

import (
	"errors"
	"fmt"
	"wkm/utils"

	"gorm.io/gorm"
)

type MerchantRepository interface {
	CreateMerchant(data Merchant) error
	MasterData(search string, limit int, pageParams int) []Merchant
	MasterDataCount(search string) int64
	DetailMerchant(id string) Merchant
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

func (lR *merchantRepository) DetailMerchant(id string) Merchant {
	merchant := Merchant{ID: id}
	lR.conn.Preload("Kategori").Preload("MediaPromosi").Preload("NamaPICMRO").Find(&merchant)
	return merchant
}

func (lR *merchantRepository) CreateMerchant(data Merchant) error {
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

func (lR *merchantRepository) Update(data Merchant) error {
	record := Merchant{ID: data.ID}
	lR.conn.First(&record)
	if record.Nama == "" {
		return errors.New("maaf data tidak ada")
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

func (lR *merchantRepository) MasterData(search string, limit int, pageParams int) []Merchant {
	merchant := []Merchant{}
	query := lR.conn.Where("nama like ? or alamat like  ? ", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&merchant)
	return merchant
}

func (lR *merchantRepository) MasterDataCount(search string) int64 {
	var merchant []Merchant
	query := lR.conn.Where("nama like ? or alamat like  ? ", "%"+search+"%", "%"+search+"%")
	query.Select("id").Find(&merchant)
	return int64(len(merchant))
}