package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func CreateSchema(pg *pg.DB) error {
	//temporary..
	var gift Gift

	err := pg.Model(&gift).CreateTable(&orm.CreateTableOptions{
		Temp:          false,
		FKConstraints: true,
	})

	return err
}
