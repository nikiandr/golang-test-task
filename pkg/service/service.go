package service

import (
	"github.com/nikiandr/golang-test-task/pkg/model"
	"github.com/nikiandr/golang-test-task/pkg/repository"
)

type Member interface {
	GetAllMembers() ([]model.Member, error)
	CreateMember(member model.MemberInput) (int, error)
	DeleteMember(memberId int) error
}

type Service struct {
	Member
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Member: NewMemberService(repos),
	}
}
