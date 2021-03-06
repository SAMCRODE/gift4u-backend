package main

import (
	"github.com/SAMCRODE/gift4u-backend/config"
	"github.com/SAMCRODE/gift4u-backend/db"
	"github.com/SAMCRODE/gift4u-backend/models"
	"github.com/SAMCRODE/gift4u-backend/server"

	"github.com/SAMCRODE/gift4u-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwapper "github.com/swaggo/gin-swagger"
)

func main() {
	config.Init()
	db.Init()

	docs.SwaggerInfo.Title = "Swagger gift4u API"
	docs.SwaggerInfo.Description = "Docs for giftu backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "gift4uall.herokuapp.com"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	models.CreateSchema(db.GetDB())

	r := server.NewRouter()

	r.GET("/swagger/*any", ginSwapper.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
