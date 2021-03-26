package db

import (
	"fmt"
	"gift4u/config"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func Init() {
	conf := config.GetConfig()

	opt, err := pg.ParseURL(conf.DATABASEURI)
	if err != nil {
		fmt.Println("Database connection error")
		// panic(err)
	} else {
		db = pg.Connect(opt)
		fmt.Println(conf.DATABASEURI)
	}
}

func GetDB() *pg.DB {
	return db
}
