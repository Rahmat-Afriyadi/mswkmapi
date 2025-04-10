package mediaPromosi

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type MediaPromosiRepository interface {
	CreateMediaPromosi(data entity.MediaPromosi) error
	MasterDataAll() []entity.MediaPromosi
	MasterData(search string, limit int, pageParams int) []entity.MediaPromosi
	MasterDataCount(search string) int64
	DetailMediaPromosi(id string) entity.MediaPromosi
	Update(body entity.MediaPromosi) error
	Delete(id string, name string) error
}

type mediaPromosiRepository struct {
	conn *gorm.DB
}

func NewMediaPromosiRepository(conn *gorm.DB) MediaPromosiRepository {
	return &mediaPromosiRepository{
		conn: conn,
	}
}

func (lR *mediaPromosiRepository) MasterDataAll() []entity.MediaPromosi {
	var mediaPromosis []entity.MediaPromosi
	lR.conn.Select("id, nama").Find(&mediaPromosis)
	return mediaPromosis
}

func (lR *mediaPromosiRepository) DetailMediaPromosi(id string) entity.MediaPromosi {
	mediaPromosi := entity.MediaPromosi{ID: id}
	lR.conn.Find(&mediaPromosi)
	return mediaPromosi
}

func (lR *mediaPromosiRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&entity.MediaPromosi{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *mediaPromosiRepository) CreateMediaPromosi(data entity.MediaPromosi) error {
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *mediaPromosiRepository) Update(data entity.MediaPromosi) error {
	var record entity.MediaPromosi
	if err := lR.conn.First(&record, data.ID).Error; err != nil {
		return errors.New("maaf data tidak ada")
	}
	if err := lR.conn.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (lR *mediaPromosiRepository) MasterData(search string, limit int, pageParams int) []entity.MediaPromosi {
	mediaPromosi := []entity.MediaPromosi{}
	query := lR.conn.Where("nama like ? ", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&mediaPromosi)
	return mediaPromosi
}

func (lR *mediaPromosiRepository) MasterDataCount(search string) int64 {
	var mediaPromosi []entity.MediaPromosi
	query := lR.conn.Where("nama like ? ", "%"+search+"%")
	query.Select("id").Find(&mediaPromosi)
	return int64(len(mediaPromosi))
}
