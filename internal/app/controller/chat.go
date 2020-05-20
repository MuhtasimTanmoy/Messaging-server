package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Chat controller
func Chat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.tmpl", gin.H{
		"title": "Messaging",
	})
}