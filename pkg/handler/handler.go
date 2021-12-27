package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikiandr/golang-test-task/pkg/model"
)

func InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(Logger())

	members_data := []model.Member{}
	router.LoadHTMLGlob("assets/html/*")
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{
			"error":   "",
			"members": members_data,
		})
	})

	router.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		unique_email := true
		for _, member := range members_data {
			if email == member.Email {
				unique_email = false
				break
			}
		}
		if unique_email {
			new_member := model.Member{
				Id:        len(members_data) + 1,
				Name:      name,
				Email:     email,
				CreatedAt: time.Now().UTC().Format("2006-01-02"),
			}
			members_data = append(members_data, new_member)
			c.HTML(http.StatusOK, "main.html", gin.H{
				"error":   "",
				"members": members_data,
			})
		} else {
			c.HTML(http.StatusOK, "main.html", gin.H{
				"error":   "This email already exists.",
				"members": members_data,
			})
		}
	})

	return router
}

/* func (h *Handler) createMember(c *gin.Context) {
	content := c.PostForm("memberForm")
	c.
} */
