package banner

import "wkm/entity"

type BannerService interface {
	MasterDataPin() ([]entity.Banner,error)
	CreateBanner(data entity.Banner) error
	MasterData(search string, limit int, pageParams int) []entity.Banner
	MasterDataCount(search string) int64
	MasterDataSearch(search string) []entity.Banner
	DetailBanner(id string) entity.Banner
	Update(body entity.Banner) error
	MasterDataAll() []entity.Banner
	Delete(id string, name string) error
}

type bannerService struct {
	trR BannerRepository
}

func NewBannerService(tR BannerRepository) BannerService {
	return &bannerService{
		trR: tR,
	}
}

func (s *bannerService) MasterDataAll() []entity.Banner {
	return s.trR.MasterDataAll()
}
func (s *bannerService) MasterDataPin() ([]entity.Banner, error) {
	return s.trR.MasterDataPin()
}
func (s *bannerService) MasterDataSearch(search string) []entity.Banner {
	return s.trR.MasterDataSearch(search)
}
func (s *bannerService) Update(body entity.Banner) error {
	return s.trR.Update(body)
}
func (s *bannerService) Delete(id string, name string) error {
	return s.trR.Delete(id, name)
}
func (s *bannerService) DetailBanner(id string) entity.Banner {
	return s.trR.DetailBanner(id)
}
func (s *bannerService) MasterData(search string, limit int, pageParams int) []entity.Banner {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *bannerService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *bannerService) CreateBanner(data entity.Banner) error {
	return s.trR.CreateBanner(data)
}
