package picMro

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type PicMroRepository interface {
	CreatePicMro(data entity.PicMro) error
	MasterDataAll() []entity.PicMro
	MasterData(search string, limit int, pageParams int) []entity.PicMro
	MasterDataCount(search string) int64
	DetailPicMro(id string) entity.PicMro
	Update(body entity.PicMro) error
	Delete(id string, name string) error
}

type picMroRepository struct {
	conn *gorm.DB
}

func NewPicMroRepository(conn *gorm.DB) PicMroRepository {
	return &picMroRepository{
		conn: conn,
	}
}

func (lR *picMroRepository) MasterDataAll() []entity.PicMro {
	var picMros []entity.PicMro
	lR.conn.Select("id, nama").Where("is_deleted = 0").Find(&picMros)
	return picMros
}

func (lR *picMroRepository) DetailPicMro(id string) entity.PicMro {
	picMro := entity.PicMro{ID: id}
	lR.conn.Find(&picMro)
	return picMro
}

func (lR *picMroRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&entity.PicMro{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *picMroRepository) CreatePicMro(data entity.PicMro) error {
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *picMroRepository) Update(data entity.PicMro) error {
	record := entity.PicMro{ID: data.ID}
	if err := lR.conn.First(&record).Error; err != nil {
		return errors.New("maaf data tidak ada")
	}
	if err := lR.conn.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (lR *picMroRepository) MasterData(search string, limit int, pageParams int) []entity.PicMro {
	picMro := []entity.PicMro{}
	query := lR.conn.Where("nama like ? ", "%"+search+"%").Where("is_deleted = 0")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&picMro)
	return picMro
}

func (lR *picMroRepository) MasterDataCount(search string) int64 {
	var picMro []entity.PicMro
	query := lR.conn.Where("nama like ? ", "%"+search+"%").Where("is_deleted = 0")
	query.Select("id").Find(&picMro)
	return int64(len(picMro))
}
