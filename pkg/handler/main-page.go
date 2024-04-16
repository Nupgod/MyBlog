package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"pagination": "pagination",
	})
}