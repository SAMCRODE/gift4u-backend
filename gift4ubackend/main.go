package main

import (
	"gift4u/config"
	"gift4u/db"
	"gift4u/models"
	"gift4u/server"
)

func main() {
	config.Init()
	db.Init()

	models.CreateSchema(db.GetDB())

	r := server.NewRouter()

	r.Run()
}
