package banner

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type BannerRepository interface {
	MasterDataPin() ([]entity.Banner,error)
	CreateBanner(data entity.Banner) error
	MasterDataAll() []entity.Banner
	MasterData(search string, limit int, pageParams int) []entity.Banner
	MasterDataCount(search string) int64
	MasterDataSearch(search string) []entity.Banner

	DetailBanner(id string) entity.Banner
	Update(body entity.Banner) error
	Delete(id string, name string) error
}

type bannerRepository struct {
	conn *gorm.DB
}

func NewBannerRepository(conn *gorm.DB) BannerRepository {
	return &bannerRepository{
		conn: conn,
	}
}

func (lR *bannerRepository) MasterDataPin() ([]entity.Banner,error) {
	banner := []entity.Banner{}
	lR.conn.Where("pin = 1 and is_deleted=0").Select("id, banner, pin").Limit(5).Find(&banner).Order("updated_at desc")
	return banner,nil
}


func (lR *bannerRepository) MasterDataAll() []entity.Banner {
	var banner []entity.Banner
	lR.conn.Select("id, nama").Where("is_deleted = 0").Find(&banner)
	return banner
}

func (lR *bannerRepository) MasterDataSearch(search string) []entity.Banner {
	banner := []entity.Banner{}
	query := lR.conn.Select("DISTINCT banner.id, banner.nama").Where("banner.is_deleted = 0")
	if search != "" {
		query.Joins("JOIN banner_kategoris a ON a.BannerID = banner.id").
			Joins("JOIN mst_kategori_banner b ON a.BannerKategoriID = b.id").
			Where("banner.nama like ? or banner.deskripsi like ? or b.nama like ? ", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	query.Limit(3).Find(&banner)
	return banner
}


func (lR *bannerRepository) DetailBanner(id string) entity.Banner {
	banner := entity.Banner{ID: id}
	lR.conn.Preload("Kategori").Find(&banner)
	return banner
}

func (lR *bannerRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&entity.Banner{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *bannerRepository) CreateBanner(data entity.Banner) error {
	var count int64
	lR.conn.Model(&entity.Banner{}).Where("pin = 1 and is_deleted=0").Count(&count)
	if count >= 12 && data.Pin {
		return errors.New("maaf banner pin sudah mencapai batas maksimal")
		
	}
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *bannerRepository) Update(data entity.Banner) error {
	record := entity.Banner{ID: data.ID}
	var count int64
	lR.conn.Model(&entity.Banner{}).Where("pin = 1 and is_deleted=0").Count(&count)
	lR.conn.First(&record)
	if record.Banner == "" {
		return errors.New("maaf data tidak ada")
	}
	if count >= 12 && !record.Pin && data.Pin {
		return errors.New("maaf banner pin sudah mencapai batas maksimal")
	}
	if err := lR.conn.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (lR *bannerRepository) MasterData(search string, limit int, pageParams int) []entity.Banner {
	banner := []entity.Banner{}
	query := lR.conn.Where("is_deleted = 0")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&banner)
	return banner
}

func (lR *bannerRepository) MasterDataCount(search string) int64 {
	var banner []entity.Banner
	query := lR.conn.Where("is_deleted = 0")
	query.Select("id").Find(&banner)
	return int64(len(banner))
}
