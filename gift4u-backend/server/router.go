package server

import (
	"github.com/SAMCRODE/gift4u-backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	giftGroup := router.Group("gift")
	{
		gift := new(controllers.GiftController)
		giftGroup.POST("/create", gift.Save)
		giftGroup.GET("/search/:id", gift.SearchById)
		giftGroup.GET("/collection/:page", gift.SearchGiftsByPage)
	}

	return router
}
