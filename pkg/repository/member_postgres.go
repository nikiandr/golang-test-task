package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/nikiandr/golang-test-task/pkg/model"
)

type MemberPostgres struct {
	db *sqlx.DB
}

func NewMemberPostgres(db *sqlx.DB) *MemberPostgres {
	return &MemberPostgres{db: db}
}

func (r *MemberPostgres) CreateMember(member model.MemberInput) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, email) VALUES ($1, $2) RETURNING id", membersTable)

	var memberId int
	err := r.db.Get(&memberId, query, member.Name, member.Email)
	if err != nil {
		return 0, err
	}

	return memberId, nil
}

func (r *MemberPostgres) GetAllMembers() ([]model.Member, error) {
	var members []model.Member

	query := fmt.Sprintf("SELECT * FROM %s", membersTable)
	err := r.db.Select(&members, query)

	if err != nil {
		return []model.Member{}, err
	}

	//Used to truncate last T...Z part
	//Example:
	//2021-12-24T00:00:00Z -> 2021-12-24
	for i := 0; i < len(members); i++ {
		members[i].CreatedAt = strings.Split(members[i].CreatedAt, "T")[0]
	}
	return members, nil
}

func (r *MemberPostgres) DeleteMember(memberId int) error {
	var returningMemberId int

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", membersTable)
	err := r.db.Get(&returningMemberId, query, memberId)
	if err != nil {
		return err
	}

	return nil
}
