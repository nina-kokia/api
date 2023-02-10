package router

import (
	"api/handlers"
	"api/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	authorized := router.Group("/")
	authorized.Use(middleware.HeaderToken())

	router.GET("/", handlers.Ready)

	authorized.POST("/api/add/:id/:ranking", handlers.Register)
	authorized.POST("/api/add/np/:coupon", handlers.RegisterCoupon)
	authorized.DELETE("/api/delete/:coupon", handlers.DeleteCoupon)
	authorized.GET("/api/find/:id", handlers.Find)
	authorized.GET("/api/list/all", handlers.FindAll)
	router.NoRoute(handlers.Err)

	fmt.Printf("Starting server in Port %v\n", middleware.Port())
	router.Run(":" + middleware.Port())
}
