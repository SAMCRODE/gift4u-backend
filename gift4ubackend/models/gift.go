package models

import (
	"gift4u/db"
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

func (g Gift) Save() error {
	pg := db.GetDB()
	_, err := pg.Model(&g).Returning("ID").Insert()
	return err
}
