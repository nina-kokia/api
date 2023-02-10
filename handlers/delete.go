package handlers

import (
	"api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCoupon(c *gin.Context) {

	err := database.Delete(c.Param("coupon"))
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
