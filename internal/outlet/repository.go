package outlet

import (
	"errors"
	"fmt"
	"wkm/utils"

	"gorm.io/gorm"
)

type OutletRepository interface {
	CreateOutlet(data Outlet) error
	MasterData(search string, limit int, pageParams int) []Outlet
	MasterDataCount(search string) int64
	DetailOutlet(id string) Outlet
	Update(body Outlet) error
}

type outletRepository struct {
	conn *gorm.DB
}

func NewOutletRepository(conn *gorm.DB) OutletRepository {
	return &outletRepository{
		conn: conn,
	}
}

func (lR *outletRepository) DetailOutlet(id string) Outlet {
	outlet := Outlet{ID: id}
	lR.conn.Preload("MediaPromosi").Preload("Merchant").Preload("Merchant.Kategori").Find(&outlet)
	fmt.Println("ini outlet yaa ", outlet)
	return outlet
}

func (lR *outletRepository) CreateOutlet(data Outlet) error {
	mediaPromosi := data.MediaPromosi
	result := lR.conn.Omit("MediaPromosi").Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	err := lR.conn.Model(&data).Association("MediaPromosi").Replace(mediaPromosi)
	if err != nil {
		return err
	}
	return nil
}

func (lR *outletRepository) Update(data Outlet) error {
	record := Outlet{ID: data.ID}
	lR.conn.Find(&record)
	mediaPromosi := data.MediaPromosi
	if record.Nama == "" {
		return errors.New("maaf data tidak ada")
	}
	result := lR.conn.Omit("MediaPromosi").Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	err := lR.conn.Model(&record).Association("MediaPromosi").Replace(mediaPromosi)
	if err != nil {
		return err
	}
	return nil
}

func (lR *outletRepository) MasterData(search string, limit int, pageParams int) []Outlet {
	outlet := []Outlet{}
	query := lR.conn.Where("nama like ? or alamat like  ? ", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&outlet)
	return outlet
}

func (lR *outletRepository) MasterDataCount(search string) int64 {
	var outlet []Outlet
	query := lR.conn.Where("nama like ? or alamat like  ? ", "%"+search+"%", "%"+search+"%")
	query.Select("id").Find(&outlet)
	return int64(len(outlet))
}
