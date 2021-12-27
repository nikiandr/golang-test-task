package handler

import (
	"net/http"
	"net/mail"
	"regexp"
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
		// Check for email uniqueness
		unique_email := true
		for _, member := range members_data {
			if email == member.Email {
				unique_email = false
				break
			}
		}
		if !unique_email {
			c.HTML(http.StatusBadRequest, "main.html", gin.H{
				"error":   "This email already exists.",
				"members": members_data,
			})
		} else {
			// Validate email format
			_, err := mail.ParseAddress(email)
			correct_email := (err == nil)
			// Validate name format
			r, _ := regexp.Compile("^[ A-Za-z.]+$")
			correct_name := r.Match([]byte(name))
			if !correct_email || !correct_name {
				c.HTML(http.StatusBadRequest, "main.html", gin.H{
					"error":   "Incorrect email or name format.",
					"members": members_data,
				})
			} else {
				locat, _ := time.LoadLocation("Europe/Kiev")
				new_member := model.Member{
					Id:        len(members_data) + 1,
					Name:      name,
					Email:     email,
					CreatedAt: time.Now().In(locat).Format("2006-01-02"),
				}
				members_data = append(members_data, new_member)
				c.HTML(http.StatusOK, "main.html", gin.H{
					"error":   "",
					"members": members_data,
				})
			}
		}
	})

	return router
}

/* func (h *Handler) createMember(c *gin.Context) {
	content := c.PostForm("memberForm")
	c.
} */
