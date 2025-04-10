package mediaPromosi

import "wkm/entity"

type MediaPromosiService interface {
	CreateMediaPromosi(data entity.MediaPromosi) error
	MasterData(search string, limit int, pageParams int) []entity.MediaPromosi
	MasterDataCount(search string) int64
	DetailMediaPromosi(id string) entity.MediaPromosi
	Update(body entity.MediaPromosi) error
	MasterDataAll() []entity.MediaPromosi
	Delete(id string, name string) error
}

type mediaPromosiService struct {
	trR MediaPromosiRepository
}

func NewMediaPromosiService(tR MediaPromosiRepository) MediaPromosiService {
	return &mediaPromosiService{
		trR: tR,
	}
}

func (s *mediaPromosiService) MasterDataAll() []entity.MediaPromosi {
	return s.trR.MasterDataAll()
}
func (s *mediaPromosiService) Update(body entity.MediaPromosi) error {
	return s.trR.Update(body)
}
func (s *mediaPromosiService) Delete(id string, name string) error {
	return s.trR.Delete(id, name)
}
func (s *mediaPromosiService) DetailMediaPromosi(id string) entity.MediaPromosi {
	return s.trR.DetailMediaPromosi(id)
}
func (s *mediaPromosiService) MasterData(search string, limit int, pageParams int) []entity.MediaPromosi {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *mediaPromosiService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *mediaPromosiService) CreateMediaPromosi(data entity.MediaPromosi) error {
	return s.trR.CreateMediaPromosi(data)
}
