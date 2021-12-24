package model

type Member struct {
	Name  string `db:"name"`
	Email string `db:"email"`
}
