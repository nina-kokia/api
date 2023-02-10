package handlers

import (
	"api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {

	request, err := database.Find(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
	}

	if request == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Anime not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"users":   request,
	})
}
