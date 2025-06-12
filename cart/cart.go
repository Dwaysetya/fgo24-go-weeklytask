package cart

import (
	"fmt"
	"sync"
)

type Item interface {
	GetName() string
	GetPrice() float64
}

type Cart struct {
	items []Item
	mu    sync.Mutex
}

func NewCart() *Cart {
	return &Cart{
		items: []Item{},
	}
}

func (c *Cart) AddItem(item Item) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = append(c.items, item)
}

func (c *Cart) PrintCart() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.items) == 0 {
		fmt.Println("Belum ada belanjaan")
		return
	}

	fmt.Println("\n🛒 Isi Keranjang:")
	total := 0.0
	for i, item := range c.items {
		fmt.Printf("%d. %s - Rp %.0f\n", i+1, item.GetName(), item.GetPrice())
		total += item.GetPrice()
	}
	fmt.Printf("Total: Rp %.0f\n", total)
}
