package handlers

import (
	"api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	err := database.Register(c.Param("id"), c.Param("ranking"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
	})
}

func RegisterCoupon(c *gin.Context) {

	err := database.RegisterCoupon(c.Param("coupon"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
	})
}
