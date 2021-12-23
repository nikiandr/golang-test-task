package club

type Member struct {
	Name  string `db:"name"`
	Email string `db:"email"`
}
