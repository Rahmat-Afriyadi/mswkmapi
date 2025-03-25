package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type MemberService interface {
	Mine(userId string) ([]entity.Member, error)
	AddCard(kode request.AddMemberCard) (entity.Member, error)
	CreateNewMemberCard(member request.CreateNewMember) (entity.Member, error)
}

type memberService struct {
	mR repository.MemberRepository
}

func NewMemberService(mR repository.MemberRepository) MemberService {
	return &memberService{
		mR,
	}
}

func (s *memberService) Mine(userId string) ([]entity.Member, error) {
	return s.mR.Mine(userId), nil
}

func (s *memberService) AddCard(data request.AddMemberCard) (entity.Member, error) {
	return s.mR.AddCard(data)
}
func (s *memberService) CreateNewMemberCard(data request.CreateNewMember) (entity.Member, error) {
	return s.mR.CreateNewMemberCard(data)
}
