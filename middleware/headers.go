package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Request.Header.Add(os.Getenv("HEADER"), os.Getenv("TOKEN"))
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		c.Next()
	}
}

func HeaderToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header = c.GetHeader(os.Getenv("HEADER"))
		header = strings.TrimSpace(header)

		if header != os.Getenv("TOKEN") {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Missing auth token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
