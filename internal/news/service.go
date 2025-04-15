package news

import "wkm/entity"

type NewsService interface {
	CreateNews(data entity.News) error
	MasterData(search string, limit int, pageParams int) []entity.News
	MasterDataCount(search string) int64
	MasterDataSearch(search string) []entity.News
	DetailNews(id string) entity.News
	Update(body entity.News) error
	MasterDataAll() []entity.News
	Delete(id string, name string) error
}

type newsService struct {
	trR NewsRepository
}

func NewNewsService(tR NewsRepository) NewsService {
	return &newsService{
		trR: tR,
	}
}

func (s *newsService) MasterDataAll() []entity.News {
	return s.trR.MasterDataAll()
}
func (s *newsService) MasterDataSearch(search string) []entity.News {
	return s.trR.MasterDataSearch(search)
}
func (s *newsService) Update(body entity.News) error {
	return s.trR.Update(body)
}
func (s *newsService) Delete(id string, name string) error {
	return s.trR.Delete(id, name)
}
func (s *newsService) DetailNews(id string) entity.News {
	return s.trR.DetailNews(id)
}
func (s *newsService) MasterData(search string, limit int, pageParams int) []entity.News {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *newsService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *newsService) CreateNews(data entity.News) error {
	return s.trR.CreateNews(data)
}
