package models

import (
	"gift4u/config"
	"gift4u/db"
	"testing"
)

func TestInsert(t *testing.T) {
	// if 1 != 1 {
	// 	t.Errorf("Abs(-1) = %d; want 1", got)
	// }
	config.Init()
	db.Init()

	var g Gift

	g.Name = "playstation 2"
	g.Description = "slim"
	g.Value = 200.00
	g.Boleto = "1234.1234.1234.1234"
	g.PixCode = "http://qrcode.com/qr123.jpg"
	g.ProductImage = "http://playstationfera.com/play.jpg"
	err2 := g.Save()

	if err2 != nil {
		// panic(err2)
		t.Errorf("err")
	}
}