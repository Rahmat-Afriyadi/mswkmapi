package repository

import (
	"errors"
	"wkm/entity"
	"wkm/request"

	"gorm.io/gorm"
)

type MemberRepository interface {
	Mine(userId string) []entity.Member
	AddCard(data request.AddMemberCard) (entity.Member, error)
}

type memberRepository struct {
	connUser *gorm.DB
}

func NewMemberRepository(connUser *gorm.DB) MemberRepository {
	return &memberRepository{
		connUser: connUser,
	}
}

func (s *memberRepository) Mine(userId string) []entity.Member {
	data := []entity.Member{}
	s.connUser.Where("user_id", userId).Find(&data)
	return data
}

func (s *memberRepository) AddCard(data request.AddMemberCard) (entity.Member, error) {
	card := entity.Member{}
	s.connUser.Where("no_msn = ? or no_kartu = ?", data.Kode, data.Kode).Where("nm_customer11 = ?", data.Nama).First(&card)
	if card.NoMsn == "" {
		return entity.Member{}, errors.New("nomor kartu atau nomor mesin tidak ditemukan")
	}
	if card.UserId != nil {
		return entity.Member{}, errors.New("nomor kartu tersebut telah digunakan oleh seseorang silahkan hubungi admin")
	}
	card.UserId = &data.UserID
	s.connUser.Save(&card)
	return card, nil
}
