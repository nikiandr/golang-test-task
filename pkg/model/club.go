package model

type Member struct {
	Id        int
	Name      string `form:"name" binding:"required"`
	Email     string `form:"email" binding:"required"`
	CreatedAt string
}

type MemberInput struct {
	Name  string
	Email string
}
