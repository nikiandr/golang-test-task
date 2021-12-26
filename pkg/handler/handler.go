package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikiandr/golang-test-task/pkg/model"
)

func InitRoutes() *gin.Engine {
	router := gin.New()

	members_data := []model.Member{}

	router.LoadHTMLGlob("assets/html/*")
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", members_data)
	})

	router.POST("/", func(c *gin.Context) {
		new_member := model.Member{
			Id:        len(members_data) + 1,
			Name:      c.PostForm("name"),
			Email:     c.PostForm("email"),
			CreatedAt: time.Now().UTC().Format("2006-01-02"),
		}
		c.BindJSON(&new_member)
		new_member.Id = len(members_data) + 1
		new_member.CreatedAt = time.Now().UTC().Format("2006-01-02")
		members_data = append(members_data, new_member)
		c.HTML(http.StatusOK, "main.html", members_data)
	})

	return router
}

/* func (h *Handler) createMember(c *gin.Context) {
	content := c.PostForm("memberForm")
	c.
} */
