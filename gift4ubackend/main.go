package main

import (
	"gift4u/config"
	"gift4u/db"
	"gift4u/models"
	"gift4u/server"

	"docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwapper "github.com/swaggo/gin-swagger"
)

func main() {
	config.Init()
	db.Init()

	docs.SwaggerInfo.Title = "Swagger gift4u API"
	docs.SwaggerInfo.Description = "Docs for giftu backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "api.gift4u.com"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	models.CreateSchema(db.GetDB())

	r := server.NewRouter()

	r.GET("/swagger/*any", ginSwapper.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
