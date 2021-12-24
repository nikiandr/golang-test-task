package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderMainPage(c *gin.Context) {
	memberList, err := h.services.GetAllMembers()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "assets/html/main.html", memberList)
}

/* func (h *Handler) createMember(c *gin.Context) {
	content := c.PostForm("memberForm")
	c.
} */
