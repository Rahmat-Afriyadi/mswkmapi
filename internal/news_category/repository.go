package newsKategori

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type NewsKategoriRepository interface {
	CreateNewsKategori(data entity.NewsKategori) error
	MasterDataAll() []entity.NewsKategori
	MasterData(search string, limit int, pageParams int) []entity.NewsKategori
	MasterDataCount(search string) int64
	DetailNewsKategori(id string) entity.NewsKategori
	Update(body entity.NewsKategori) error
	Delete(id string, name string) error
}

type kategoriRepository struct {
	conn *gorm.DB
}

func NewNewsKategoriRepository(conn *gorm.DB) NewsKategoriRepository {
	return &kategoriRepository{
		conn: conn,
	}
}

func (lR *kategoriRepository) MasterDataAll() []entity.NewsKategori {
	var kategoris []entity.NewsKategori
	lR.conn.Select("id, nama").Where("is_deleted = 0").Find(&kategoris)
	return kategoris
}

func (lR *kategoriRepository) DetailNewsKategori(id string) entity.NewsKategori {
	kategori := entity.NewsKategori{ID: id}
	lR.conn.Find(&kategori)
	return kategori
}

func (lR *kategoriRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&entity.NewsKategori{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *kategoriRepository) CreateNewsKategori(data entity.NewsKategori) error {
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *kategoriRepository) Update(data entity.NewsKategori) error {
	record := entity.NewsKategori{ID: data.ID}
	if err := lR.conn.First(&record).Error; err != nil {
		return errors.New("maaf data tidak ada")
	}
	if err := lR.conn.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (lR *kategoriRepository) MasterData(search string, limit int, pageParams int) []entity.NewsKategori {
	kategori := []entity.NewsKategori{}
	query := lR.conn.Where("nama like ? ", "%"+search+"%").Where("is_deleted = 0")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&kategori)
	return kategori
}

func (lR *kategoriRepository) MasterDataCount(search string) int64 {
	var kategori []entity.NewsKategori
	query := lR.conn.Where("nama like ? ", "%"+search+"%").Where("is_deleted = 0")
	query.Select("id").Find(&kategori)
	return int64(len(kategori))
}
