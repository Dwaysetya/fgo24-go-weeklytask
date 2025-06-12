package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"weeklytask/cart"
	"weeklytask/menu"
	"weeklytask/utils"
)

func main() {
	defer fmt.Println("Terima kasih telah berkunjung")

	shoppingCart := cart.NewCart()

	for {
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Pesan Makanan")
		fmt.Println("2. Pesan Minuman")
		fmt.Println("3. Pesan Snack")
		fmt.Println("4. Lihat Keranjang")
		fmt.Println("5. Checkout")

		choice := utils.ReadInput("Pilih menu (1-5): ")

		switch choice {
		case 1:
			menu.Order(menu.Foods, shoppingCart)
		case 2:
			menu.Order(menu.Drinks, shoppingCart)
		case 3:
			menu.Order(menu.Snacks, shoppingCart)
		case 4:
			shoppingCart.PrintCart()
		case 5:
			checkout(shoppingCart)
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func checkout(c *cart.Cart) {
	var wg sync.WaitGroup
	resultChan := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		c.PrintCart()
		resultChan <- "Checkout berhasil!"
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for msg := range resultChan {
		fmt.Println(msg)
	}
}
