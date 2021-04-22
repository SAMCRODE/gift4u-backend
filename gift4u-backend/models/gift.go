package models

import (
	"github.com/SAMCRODE/gift4u-backend/db"
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

func SearchGiftById(id int) (*Gift, error) {
	pg := db.GetDB()
	g := &Gift{ID: id}
	err := pg.Model(g).WherePK().Select()

	return g, err
}

func SearchGiftsPaginated(page int) ([]Gift, error) {
	pg := db.GetDB()
	var gifts []Gift

	err := pg.Model(&gifts).Offset(10 * (page - 1)).Limit(10).Select()

	return gifts, err
}

func GetTotalGift() (int, error) {
	pg := db.GetDB()
	var gifts []Gift
	total, err := pg.Model(&gifts).Count()

	return total, err
}
