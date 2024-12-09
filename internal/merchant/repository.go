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
	lR.conn.Find(&merchant)
	return merchant
}

func (lR *merchantRepository) CreateMerchant(data Merchant) error {
	result := lR.conn.Create(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}

}

func (lR *merchantRepository) Update(data Merchant) error {
	record := Merchant{ID: data.ID}
	lR.conn.Find(&record)
	if record.Nama == "" {
		return errors.New("maaf data tidak ada")
	}
	result := lR.conn.Save(&data)
	// lR.Model(&user).Association("Roles").Replace(newRoles)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}
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
