package kategori

import "wkm/entity"

type KategoriService interface {
	CreateKategori(data entity.Kategori) error
	MasterData(search string, limit int, pageParams int) []entity.Kategori
	MasterDataCount(search string) int64
	DetailKategori(id string) entity.Kategori
	Update(body entity.Kategori) error
	MasterDataAll() []entity.Kategori
	Delete(id string, name string) error
}

type kategoriService struct {
	trR KategoriRepository
}

func NewKategoriService(tR KategoriRepository) KategoriService {
	return &kategoriService{
		trR: tR,
	}
}

func (s *kategoriService) MasterDataAll() []entity.Kategori {
	return s.trR.MasterDataAll()
}
func (s *kategoriService) Update(body entity.Kategori) error {
	return s.trR.Update(body)
}
func (s *kategoriService) Delete(id string, name string) error {
	return s.trR.Delete(id, name)
}
func (s *kategoriService) DetailKategori(id string) entity.Kategori {
	return s.trR.DetailKategori(id)
}
func (s *kategoriService) MasterData(search string, limit int, pageParams int) []entity.Kategori {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *kategoriService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *kategoriService) CreateKategori(data entity.Kategori) error {
	return s.trR.CreateKategori(data)
}
