package handlers

import (
	"api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {

	request, err := database.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
	}

	if request == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Coupons not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"coupons": request,
	})
}
