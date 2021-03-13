package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Gift struct {
	ID           int
	Name         string
	Description  string
	Value        float64
	Boleto       string
	PixCode      string
	ProductImage string
}

func Insert(pg *pg.DB, gift Gift) error {
	//temporary..
	pg.Model(&gift).CreateTable(&orm.CreateTableOptions{
		Temp:          false,
		FKConstraints: true,
	})

	_, err := pg.Model(&gift).Returning("ID").Insert()
	return err
}
