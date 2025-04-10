package kategori

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type KategoriRepository interface {
	CreateKategori(data entity.Kategori) error
	MasterDataAll() []entity.Kategori
	MasterData(search string, limit int, pageParams int) []entity.Kategori
	MasterDataCount(search string) int64
	DetailKategori(id string) entity.Kategori
	Update(body entity.Kategori) error
	Delete(id string, name string) error
}

type kategoriRepository struct {
	conn *gorm.DB
}

func NewKategoriRepository(conn *gorm.DB) KategoriRepository {
	return &kategoriRepository{
		conn: conn,
	}
}

func (lR *kategoriRepository) MasterDataAll() []entity.Kategori {
	var kategoris []entity.Kategori
	lR.conn.Select("id, nama").Find(&kategoris)
	return kategoris
}

func (lR *kategoriRepository) DetailKategori(id string) entity.Kategori {
	kategori := entity.Kategori{ID: id}
	lR.conn.Find(&kategori)
	return kategori
}

func (lR *kategoriRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&entity.Kategori{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *kategoriRepository) CreateKategori(data entity.Kategori) error {
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *kategoriRepository) Update(data entity.Kategori) error {
	var record entity.Kategori
	if err := lR.conn.First(&record, data.ID).Error; err != nil {
		return errors.New("maaf data tidak ada")
	}
	if err := lR.conn.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (lR *kategoriRepository) MasterData(search string, limit int, pageParams int) []entity.Kategori {
	kategori := []entity.Kategori{}
	query := lR.conn.Where("nama like ? ", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&kategori)
	return kategori
}

func (lR *kategoriRepository) MasterDataCount(search string) int64 {
	var kategori []entity.Kategori
	query := lR.conn.Where("nama like ? ", "%"+search+"%")
	query.Select("id").Find(&kategori)
	return int64(len(kategori))
}
