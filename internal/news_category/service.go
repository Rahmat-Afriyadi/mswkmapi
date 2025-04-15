package newsKategori

import "wkm/entity"

type NewsKategoriService interface {
	CreateNewsKategori(data entity.NewsKategori) error
	MasterData(search string, limit int, pageParams int) []entity.NewsKategori
	MasterDataCount(search string) int64
	DetailNewsKategori(id string) entity.NewsKategori
	Update(body entity.NewsKategori) error
	MasterDataAll() []entity.NewsKategori
	Delete(id string, name string) error
}

type kategoriService struct {
	trR NewsKategoriRepository
}

func NewNewsKategoriService(tR NewsKategoriRepository) NewsKategoriService {
	return &kategoriService{
		trR: tR,
	}
}

func (s *kategoriService) MasterDataAll() []entity.NewsKategori {
	return s.trR.MasterDataAll()
}
func (s *kategoriService) Update(body entity.NewsKategori) error {
	return s.trR.Update(body)
}
func (s *kategoriService) Delete(id string, name string) error {
	return s.trR.Delete(id, name)
}
func (s *kategoriService) DetailNewsKategori(id string) entity.NewsKategori {
	return s.trR.DetailNewsKategori(id)
}
func (s *kategoriService) MasterData(search string, limit int, pageParams int) []entity.NewsKategori {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *kategoriService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *kategoriService) CreateNewsKategori(data entity.NewsKategori) error {
	return s.trR.CreateNewsKategori(data)
}
