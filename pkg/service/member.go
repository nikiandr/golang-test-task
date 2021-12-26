package service

import (
	"github.com/nikiandr/golang-test-task/pkg/model"
	"github.com/nikiandr/golang-test-task/pkg/repository"
)

type MemberService struct {
	repo repository.Member
}

func NewMemberService(repo repository.Member) *MemberService {
	return &MemberService{repo: repo}
}

func (s *MemberService) GetAllMembers() ([]model.Member, error) {
	return s.repo.GetAllMembers()
}

func (s *MemberService) CreateMember(member model.MemberInput) (int, error) {
	return s.repo.CreateMember(member)
}

func (s *MemberService) DeleteMember(memberId int) error {
	return s.repo.DeleteMember(memberId)
}
