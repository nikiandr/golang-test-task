package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nikiandr/golang-test-task/pkg/model"
)

type Member interface {
	GetAllMembers() ([]model.Member, error)
	CreateMember(member model.MemberInput) (int, error)
	DeleteMember(memberId int) error
}

type Repository struct {
	Member
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Member: NewMemberPostgres(db),
	}
}
