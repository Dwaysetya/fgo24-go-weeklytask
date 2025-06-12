package menu

import (
	"fmt"
	"weeklytask/cart"
	"weeklytask/utils"
)

type Item struct {
	Name  string
	Price float64
}

type MenuItem interface {
	GetName() string
	GetPrice() float64
}

func (i *Item) GetName() string  { return i.Name }
func (i *Item) GetPrice() float64 { return i.Price }

var Foods = []MenuItem{
	&Item{"Nasi Goreng", 15000},
	&Item{"Mie Ayam", 12000},
	&Item{"Sate Ayam", 18000},
}

var Drinks = []MenuItem{
	&Item{"Es Teh", 5000},
	&Item{"Jus Alpukat", 10000},
	&Item{"Kopi", 8000},
}

var Snacks = []MenuItem{
	&Item{"Keripik", 7000},
	&Item{"Donat", 9000},
	&Item{"Roti Bakar", 11000},
}

func Order(items []MenuItem, c *cart.Cart) {
	fmt.Println("\nPilih item:")
	for i, item := range items {
		fmt.Printf("%d. %s - Rp %.0f\n", i+1, item.GetName(), item.GetPrice())
	}
	choice := utils.ReadInput("Masukkan pilihan: ")
	if choice < 1 || choice > len(items) {
		fmt.Println("Pilihan tidak tersedia.")
		return
	}
	selected := items[choice-1]
	c.AddItem(selected)
	fmt.Println("✅ Item ditambahkan ke keranjang.")
}
