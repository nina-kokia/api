package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Hello, World!",
	})
}
