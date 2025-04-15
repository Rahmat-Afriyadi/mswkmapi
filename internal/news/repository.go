package news

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type NewsRepository interface {
	CreateNews(data entity.News) error
	MasterDataAll() []entity.News
	MasterData(search string, limit int, pageParams int) []entity.News
	MasterDataCount(search string) int64
	MasterDataSearch(search string) []entity.News

	DetailNews(id string) entity.News
	Update(body entity.News) error
	Delete(id string, name string) error
}

type newsRepository struct {
	conn *gorm.DB
}

func NewNewsRepository(conn *gorm.DB) NewsRepository {
	return &newsRepository{
		conn: conn,
	}
}

func (lR *newsRepository) MasterDataAll() []entity.News {
	var news []entity.News
	lR.conn.Select("id, nama").Where("is_deleted = 0").Find(&news)
	return news
}

func (lR *newsRepository) MasterDataSearch(search string) []entity.News {
	news := []entity.News{}
	query := lR.conn.Select("DISTINCT news.id, news.nama").Where("news.is_deleted = 0")
	if search != "" {
		query.Joins("JOIN news_kategoris a ON a.NewsID = news.id").
			Joins("JOIN mst_kategori_news b ON a.NewsKategoriID = b.id").
			Where("news.nama like ? or news.deskripsi like ? or b.nama like ? ", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	query.Limit(3).Find(&news)
	return news
}


func (lR *newsRepository) DetailNews(id string) entity.News {
	news := entity.News{ID: id}
	lR.conn.Preload("Kategori").Find(&news)
	return news
}

func (lR *newsRepository) Delete(id string, name string) error {
	result := lR.conn.Model(&entity.News{}).Where("id = ?", id).Updates(map[string]interface{}{
        "is_deleted": true,
    })
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *newsRepository) CreateNews(data entity.News) error {
	var count int64
	lR.conn.Model(&entity.News{}).Where("pin = 1 and is_deleted=0").Count(&count)
	if count >= 12 && data.Pin {
		return errors.New("maaf news pin sudah mencapai batas maksimal")
		
	}
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	return nil
}

func (lR *newsRepository) Update(data entity.News) error {
	record := entity.News{ID: data.ID}
	var count int64
	lR.conn.Model(&entity.News{}).Where("pin = 1 and is_deleted=0").Count(&count)
	lR.conn.First(&record)
	if record.Nama == "" {
		return errors.New("maaf data tidak ada")
	}
	if count >= 12 && !record.Pin && data.Pin {
		return errors.New("maaf news pin sudah mencapai batas maksimal")
	}
	if err := lR.conn.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (lR *newsRepository) MasterData(search string, limit int, pageParams int) []entity.News {
	news := []entity.News{}
	query := lR.conn.Where("nama like ? ", "%"+search+"%").Where("is_deleted = 0")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&news)
	return news
}

func (lR *newsRepository) MasterDataCount(search string) int64 {
	var news []entity.News
	query := lR.conn.Where("nama like ? ", "%"+search+"%").Where("is_deleted = 0")
	query.Select("id").Find(&news)
	return int64(len(news))
}
