package service

import (
	"wkm/entity"
	"wkm/repository"
)

type MasterDataService interface {
	KategoriMerchantAll() ([]entity.Kategori, error)
	MediaPromosiAll() ([]entity.MediaPromosi, error)
	PicMroAll() ([]entity.PicMro, error)
}

type masterDataService struct {
	mR repository.MasterDataRepository
}

func NewMasterDataService(mR repository.MasterDataRepository) MasterDataService {
	return &masterDataService{
		mR,
	}
}

func (s *masterDataService) KategoriMerchantAll() ([]entity.Kategori, error) {
	return s.mR.KategoriMerchantAll()
}
func (s *masterDataService) MediaPromosiAll() ([]entity.MediaPromosi, error) {
	return s.mR.MediaPromosiAll()
}
func (s *masterDataService) PicMroAll() ([]entity.PicMro, error) {
	return s.mR.PicMroAll()
}
