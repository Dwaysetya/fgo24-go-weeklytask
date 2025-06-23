package menu

import (
	"fmt"
	"weeklytask/cart"
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
			TampilkanKategori()
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



