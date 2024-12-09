package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type MasterDataRepository interface {
	KategoriMerchantAll() ([]entity.Kategori, error)
	MediaPromosiAll() ([]entity.MediaPromosi, error)
	PicMroAll() ([]entity.PicMro, error)
}

type masterDataRepository struct {
	connGorm *gorm.DB
}

func NewMasterDataRepository(connGorm *gorm.DB) MasterDataRepository {
	return &masterDataRepository{
		connGorm: connGorm,
	}
}

func (s *masterDataRepository) KategoriMerchantAll() ([]entity.Kategori, error) {
	datas := []entity.Kategori{}
	s.connGorm.Find(&datas)
	return datas, nil
}
func (s *masterDataRepository) MediaPromosiAll() ([]entity.MediaPromosi, error) {
	datas := []entity.MediaPromosi{}
	s.connGorm.Find(&datas)
	return datas, nil
}
func (s *masterDataRepository) PicMroAll() ([]entity.PicMro, error) {
	datas := []entity.PicMro{}
	s.connGorm.Find(&datas)
	return datas, nil
}
