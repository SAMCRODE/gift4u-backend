package models

import (
	"testing"

	"github.com/go-pg/pg/v10"
)

func TestInsert(t *testing.T) {
	// if 1 != 1 {
	// 	t.Errorf("Abs(-1) = %d; want 1", got)
	// }

	opt, err := pg.ParseURL("")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	var g Gift

	g.Name = "playstation 2"
	g.Description = "slim"
	g.Value = 200.00
	g.Boleto = "1234.1234.1234.1234"
	g.PixCode = "http://qrcode.com/qr123.jpg"
	g.ProductImage = "http://playstationfera.com/play.jpg"
	err2 := Insert(db, g)

	if err2 != nil {
		panic(err2)
		t.Errorf("err")
	}
}
