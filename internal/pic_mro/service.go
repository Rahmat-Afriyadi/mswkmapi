package picMro

import "wkm/entity"

type PicMroService interface {
	CreatePicMro(data entity.PicMro) error
	MasterData(search string, limit int, pageParams int) []entity.PicMro
	MasterDataCount(search string) int64
	DetailPicMro(id string) entity.PicMro
	Update(body entity.PicMro) error
	MasterDataAll() []entity.PicMro
	Delete(id string, name string) error
}

type picMroService struct {
	trR PicMroRepository
}

func NewPicMroService(tR PicMroRepository) PicMroService {
	return &picMroService{
		trR: tR,
	}
}

func (s *picMroService) MasterDataAll() []entity.PicMro {
	return s.trR.MasterDataAll()
}
func (s *picMroService) Update(body entity.PicMro) error {
	return s.trR.Update(body)
}
func (s *picMroService) Delete(id string, name string) error {
	return s.trR.Delete(id, name)
}
func (s *picMroService) DetailPicMro(id string) entity.PicMro {
	return s.trR.DetailPicMro(id)
}
func (s *picMroService) MasterData(search string, limit int, pageParams int) []entity.PicMro {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *picMroService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *picMroService) CreatePicMro(data entity.PicMro) error {
	return s.trR.CreatePicMro(data)
}
