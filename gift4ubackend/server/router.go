package server

import (
	"gift4u/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	giftGroup := router.Group("gift")
	{
		gift := new(controllers.GiftController)
		giftGroup.POST("/create", gift.Save)
	}

	return router
}
