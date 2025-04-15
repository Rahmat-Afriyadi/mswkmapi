package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type MasterDataRepository interface {
	KategoriMerchantAll() ([]entity.Kategori, error)
	NewsKategoriAll() ([]entity.NewsKategori, error)
	MediaPromosiAll() ([]entity.MediaPromosi, error)
	PicMroAll() ([]entity.PicMro, error)
	KodeposAll() ([]entity.Kodepos, error)
}

type masterDataRepository struct {
	connGorm *gorm.DB
}

func NewMasterDataRepository(connGorm *gorm.DB) MasterDataRepository {
	return &masterDataRepository{
		connGorm: connGorm,
	}
}

func (s *masterDataRepository) KodeposAll() ([]entity.Kodepos, error) {
	datas := []entity.Kodepos{}
	s.connGorm.Where("is_deleted = 0").Find(&datas)
	return datas, nil
}
func (s *masterDataRepository) KategoriMerchantAll() ([]entity.Kategori, error) {
	datas := []entity.Kategori{}
	s.connGorm.Where("is_deleted = 0").Find(&datas)
	return datas, nil
}
func (s *masterDataRepository) NewsKategoriAll() ([]entity.NewsKategori, error) {
	datas := []entity.NewsKategori{}
	s.connGorm.Where("is_deleted = 0").Find(&datas)
	return datas, nil
}
func (s *masterDataRepository) MediaPromosiAll() ([]entity.MediaPromosi, error) {
	datas := []entity.MediaPromosi{}
	s.connGorm.Where("is_deleted = 0").Find(&datas)
	return datas, nil
}
func (s *masterDataRepository) PicMroAll() ([]entity.PicMro, error) {
	datas := []entity.PicMro{}
	s.connGorm.Where("is_deleted = 0").Find(&datas)
	return datas, nil
}
