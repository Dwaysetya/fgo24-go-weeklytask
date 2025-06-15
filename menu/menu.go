package menu

import (
	"fmt"
	"weeklytask/cart"
	"weeklytask/models"
	"weeklytask/search"
	"weeklytask/utils"
)

func TampilkanMenu(){
	for{
		utils.ClearTerminal()
		fmt.Println(`
========== Menu Utama ==========
1. Order Menu
2. Lihat Keranjang
3. Checkout
4. search
5. Filter
0. Keluar
		`)

		var pilihan int
		fmt.Print("Silahkan Masukkan Pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan{
		case 1:
			TampilkanMenuOrder()
		case 2:
			total := cart.LihatKeranjang()
			if total > 0 {
				fmt.Print("Ingin langsung checkout? (y/n): ")
				var opsi string
				fmt.Scanln(&opsi)
				if opsi == "y" || opsi == "Y" {
					cart.Checkout()
				}
			}	
		case 3:
			cart.Checkout()
		case 4:
			search.SearchMenu()
		case 5:
			search.TampilFilterMenu()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi kami")
			return		
		default:
			fmt.Println("Pilihan tidak tersedia")
			fmt.Print("Tekan ENTER untuk kembali...")
			fmt.Scanln()

		}
	}
}

func TampilkanMenuOrder(){
	for{
		utils.ClearTerminal()
		fmt.Println(`
=============== Menu Utama ==============
1. Menu Makanan
2. Menu Minuman
3. Menu Snack
0. Kembali
		`)

		var pilihan int
		fmt.Print("Silahkan Masukkan Pilihan anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			TampilMenuMakanan()
		case 2:
			TampilMenuMinuman()
		case 3:
			TampilMenuSnack()
		case 0:
			return		
		default:
			fmt.Println("Pilihan tidak tersedia")
			fmt.Print("Tekan ENTER untuk kembali...")
			fmt.Scanln()
		}

	}
}

func TampilMenuMakanan() {
	makanan := models.GetAllMakanan()

	for {
		utils.ClearTerminal()
		item := utils.TampilkanPagination("🍽️ Menu Makanan", makanan, 5)

		if item == nil {
			return
		}

		cart.TambahkanKeranjang(item)
		fmt.Printf("✅ %s berhasil ditambahkan ke keranjang\n", item.GetName())
		fmt.Print("Tekan ENTER untuk kembali ...")
		fmt.Scanln()
	}
}


func TampilMenuMinuman() {
	minuman := models.GetAllMinuman()

	for {
		utils.ClearTerminal()
		item := utils.TampilkanPagination("🥤 Menu Makanan", minuman, 5)

		if item == nil {
			return
		}

		cart.TambahkanKeranjang(item)
		fmt.Printf("✅ %s berhasil ditambahkan ke keranjang\n", item.GetName())
		fmt.Print("Tekan ENTER untuk kembali ...")
		fmt.Scanln()
	}
}

func TampilMenuSnack() {
	snack := models.GetAllSnack()

	for {
		utils.ClearTerminal()
		item := utils.TampilkanPagination("🍿 Menu Makanan", snack, 5)

		if item == nil {
			return
		}

		cart.TambahkanKeranjang(item)
		fmt.Printf("✅ %s berhasil ditambahkan ke keranjang\n", item.GetName())
		fmt.Print("Tekan ENTER untuk kembali ...")
		fmt.Scanln()
	}

}

