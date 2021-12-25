package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikiandr/golang-test-task/pkg/model"
)

func InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("assets/html/*")
	router.GET("/", func(c *gin.Context) {
		memberList := []model.Member{}
		c.HTML(http.StatusOK, "main.html", memberList)
		memberList = append(memberList, model.Member{
			Id:        len(memberList) + 1,
			Name:      data["name"],
			Email:     data["email"],
			CreatedAt: time.Now().Format("2006-01-02"),
		})
	})

	return router
}
