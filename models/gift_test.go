package models

import (
	"fmt"
	"gift4u/db"
	"testing"

	"github.com/gift4ubackend/config"
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

func TestFindGift(t *testing.T) {
	// if 1 != 1 {
	// 	t.Errorf("Abs(-1) = %d; want 1", got)
	// }
	config.Init()
	db.Init()

	_, err := SearchGiftById(1)

	if err != nil {
		// panic(err2)
		t.Errorf("err")
	}
}

func TestSearchPaginatedGifts(t *testing.T) {
	config.Init()

	db.Init()

	gifts, err := SearchGiftsPaginated(1)

	fmt.Println(gifts)

	if err != nil {
		t.Errorf("err")
	}
}
