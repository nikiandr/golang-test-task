package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nikiandr/golang-test-task/pkg/model"
)

type MemberPostgres struct {
	db *sqlx.DB
}

func NewMemberPostgres(db *sqlx.DB) *MemberPostgres {
	return &MemberPostgres{db: db}
}

func (r *MemberPostgres) CreateMember(member model.Member) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, email) VALUES ($1, $2) RETURNING id", membersTable)

	var memberId int
	err := r.db.Get(&memberId, query, member.Name, member.Email)
	if err != nil {
		return 0, err
	}

	return memberId, nil
}

func (r *MemberPostgres) GetAllMembers() ([]model.Member, error) {
	return nil, nil
}

func (r *MemberPostgres) DeleteMember(memberId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", membersTable)
	return nil
}
