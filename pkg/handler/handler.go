package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikiandr/golang-test-task/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("assets/html/*")
	router.Static("/css", "../assets/css")
	router.GET("/", h.renderMainPage)
	// router.POST("/add", h.createMember)

	return router
}
