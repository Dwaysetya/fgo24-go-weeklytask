package menu

import (
	"fmt"
	"weeklytask/utils"
)

func TampilkanKategori(){
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