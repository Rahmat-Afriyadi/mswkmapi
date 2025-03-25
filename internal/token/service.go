package token

import "wkm/entity"

type TokenService interface {
	CreateToken(data entity.Token) error
	DetailToken(id string) (entity.Token, error)
	FindUser(id string) (entity.User, error)
	ActivateUser(token entity.Token) (entity.User, error)
}

type mstMtrService struct {
	trR TokenRepository
}

func NewTokenService(tR TokenRepository) TokenService {
	return &mstMtrService{
		trR: tR,
	}
}

func (s *mstMtrService) ActivateUser(token entity.Token) (entity.User, error) {
	return s.trR.ActivateUser(token)
}
func (s *mstMtrService) DetailToken(id string) (entity.Token, error) {
	return s.trR.DetailToken(id)
}
func (s *mstMtrService) FindUser(id string) (entity.User, error) {
	return s.trR.FindUser(id)
}

func (s *mstMtrService) CreateToken(data entity.Token) error {
	return s.trR.CreateToken(data)
}
