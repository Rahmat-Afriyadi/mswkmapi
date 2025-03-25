package repository

import (
	"errors"
	"time"
	"wkm/entity"
	"wkm/request"

	"gorm.io/gorm"
)

type MemberRepository interface {
	Mine(userId string) []entity.Member
	AddCard(data request.AddMemberCard) (entity.Member, error)
	CreateNewMemberCard(data request.CreateNewMember) (entity.Member, error)
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
	var count int64
	s.connUser.Model(&entity.Member{}).Where("no_msn = ? or no_kartu = ?", data.Kode, data.Kode).Count(&count)
	if count < 1 {
		return entity.Member{}, errors.New("nomor kartu atau nomor mesin tidak ditemukan")
	}
	s.connUser.Model(&entity.Member{}).Where("(no_msn = ? or no_kartu = ?) and user_id is null", data.Kode, data.Kode).Count(&count)
	if count < 1 {
		return entity.Member{}, errors.New("nomor kartu tersebut telah digunakan oleh seseorang silahkan hubungi admin")
	}
	s.connUser.Exec("UPDATE member SET user_id = ? WHERE (no_msn = ? or no_kartu = ?) and user_id is null", data.UserID, data.Kode, data.Kode)

	return card, nil
}

func (s *memberRepository) CreateNewMemberCard(data request.CreateNewMember) (entity.Member, error) {
	tglExpired, _ := time.Parse("2006-01-02", data.TglExpired)
	newMember := entity.Member{NoMsn: data.NoMsn, NoKartu: data.NoKartu, TglExpired: tglExpired, NmCustomer: data.NmCustomer }
	result := s.connUser.Save(&newMember)
	if result.Error != nil {
		return entity.Member{}, result.Error
	}

	return newMember, nil
}
