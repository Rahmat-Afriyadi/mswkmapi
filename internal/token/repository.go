package token

import (
	"errors"
	"wkm/entity"

	"gorm.io/gorm"
)

type TokenRepository interface {
	CreateToken(data entity.Token) error
	DetailToken(id string) (entity.Token, error)
	FindUser(id string) (entity.User, error)
	ActivateUser(token entity.Token) (entity.User, error)
}

type tokenRepository struct {
	conn *gorm.DB
}

func NewTokenRepository(conn *gorm.DB) TokenRepository {
	return &tokenRepository{
		conn: conn,
	}
}

func (lR *tokenRepository) DetailToken(id string) (entity.Token, error) {
	var token entity.Token
	lR.conn.Where("token =? and used = 0", id).First(&token)
	if token.ID == "" {
		return token, errors.New("token tidak ditemukan atau telah expired")
	}
	return token, nil
}

func (lR *tokenRepository) FindUser(id string) (entity.User, error) {
	var user entity.User
	lR.conn.Where("no_hp =?", id).First(&user)
	if user.ID == "" {
		return user, errors.New("user tidak ditemukan atau telah expired")
	}
	return user, nil
}

func (lR *tokenRepository) ActivateUser(token entity.Token) (entity.User, error) {
	var user entity.User
	lR.conn.Where("no_hp =?", token.NoHp).First(&user)
	user.Active = true
	lR.conn.Save(&user)
	token.Used = true
	lR.conn.Save(&token)
	return user, nil
}

func (lR *tokenRepository) CreateToken(data entity.Token) error {
	lR.conn.Save(&data)
	return nil
}

