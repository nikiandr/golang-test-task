package model

type Member struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	CreatedAt string `db:"created_at"`
}

type MemberInput struct {
	Name  string
	Email string
}
