package main

import (
	"gift4u/config"
	"gift4u/db"
	"gift4u/server"
)

func main() {
	config.Init()
	db.Init()

	r := server.NewRouter()

	r.Run()
}
