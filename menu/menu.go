package menu

import (
	"fmt"
	"weeklytask/cart"
	"weeklytask/models"
	"weeklytask/utils"
)

func TampilkanMenu(){
	for{
		utils.ClearTerminal()
		fmt.Println(`
=============== Menu Utama ==============
1. Menu Makanan
2. Menu Minuman
3. Menu Snack
4. Lihat Keranjang
5. Checkout
6. Keluar
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
		case 4:
			total := cart.LihatKeranjang()
			if total > 0 {
				fmt.Print("Ingin langsung checkout? (y/n): ")
				var opsi string
				fmt.Scanln(&opsi)
				if opsi == "y" || opsi == "Y" {
					cart.Checkout()
				}
			}		
		case 5:
			cart.Checkout()
		case 6:
			fmt.Println("Terima kasih telah menggunakan aplikasi kami.")
			return		
		default:
			fmt.Println("Pilihan tidak tersedia")
			fmt.Print("Tekan ENTER untuk kembali...")
			fmt.Scanln()
		}

	}
}

func TampilMenuMakanan(){
	for{
		utils.ClearTerminal()
		fmt.Println(`
1. Nasi Goreng 			Rp. 18.000
2. Nasi Padang			Rp. 15.000
3. Nasi Pecel			Rp. 12.000
4. Nasi Rames			Rp. 13.000
5. Nasi Ayam 			Rp. 17.000
6. Nasi Telur			Rp. 10.000
7. Nasi Bakar			Rp. 12.000
8. Bakso				Rp. 14.000
9. Mie Ayam				Rp. 13.000
10. Sate Ayam			Rp. 19.000

0. Kembali
		`)

		var pilihan int
		fmt.Print("Pilih menu makanan (1-10, 0 untuk kembali ): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			items:= models. NewMakanan("Nasi Goreng", 18000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅  Nasi goreng Berhasil di tambahkan")

		case 2:
			items := models. NewMakanan("Nasi Padang", 15000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅  Nasi Padang Berhasil di tambahkan")
		case 3:
			items := models. NewMakanan("Nasi Pecel", 12000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅  Nasi Pecel Berhasil di tambahkan")
		case 4:
			items := models. NewMakanan("Nasi Rames", 13000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅ Nasi Rames Berhasil di tambahkan")
		case 5:
			items := models. NewMakanan("Nasi Ayam", 17000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅ Nasi Ayam Berhasil di tambahkan")
		case 6:
			items := models. NewMakanan("Nasi Telur", 10000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅ Nasi Telur Berhasil di tambahkan")
		case 7:
			items := models. NewMakanan("Nasi Bakar", 12000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅ Nasi Bakar Berhasil di tambahkan")
		case 8:
			items := models. NewMakanan("Bakso", 14000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅ Bakso Berhasil di tambahkan")
		case 9:
			items := models. NewMakanan("Mie Ayam", 13000)
			cart.TambahkanKeranjang(items)		
			fmt.Println("✅ Mie Ayam Berhasil di tambahkan")
		case 10:
			items := models. NewMakanan("Sate Ayam", 19000)
			cart.TambahkanKeranjang(items)
			fmt.Println("✅ Sate Ayam Berhasil di tambahkan")
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Print("Tekan ENTER untuk kembali...")
		fmt.Scanln()
		return
	}
}

func TampilMenuMinuman() {
	for {
		utils.ClearTerminal()
		fmt.Println(`
1. Es Teh Manis         Rp. 5.000
2. Es Jeruk             Rp. 6.000
3. Kopi Hitam           Rp. 7.000
4. Kopi Susu            Rp. 8.000
5. Air Mineral          Rp. 4.000

0. Kembali
		`)

		var pilihan int
		fmt.Print("Pilih menu minuman (1-5, 0 untuk kembali ): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			item := models.NewMinuman("Es Teh Manis", 5000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Es Teh Manis berhasil ditambahkan")

		case 2:
			item := models.NewMinuman("Es Jeruk", 6000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Es Jeruk berhasil ditambahkan")

		case 3:
			item := models.NewMinuman("Kopi Hitam", 7000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Kopi Hitam berhasil ditambahkan")

		case 4:
			item := models.NewMinuman("Kopi Susu", 8000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Kopi Susu berhasil ditambahkan")

		case 5:
			item := models.NewMinuman("Air Mineral", 4000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Air Mineral berhasil ditambahkan")

		case 0:
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Print("Tekan ENTER untuk kembali...")
		fmt.Scanln()
		return
	}
}

func TampilMenuSnack() {
	for {
		utils.ClearTerminal()
		fmt.Println(`
1. Keripik Singkong     Rp. 5.000
2. Roti Bakar           Rp. 8.000
3. Pisang Goreng        Rp. 6.000
4. Cireng               Rp. 7.000
5. Seblak               Rp. 10.000
0. Kembali
		`)

		var pilihan int
		fmt.Print("Pilih menu snack (1-5, 0 untuk kembali): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			item := models.NewSnack("Keripik Singkong", 5000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Keripik Singkong berhasil ditambahkan")
		case 2:
			item := models.NewSnack("Roti Bakar", 8000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Roti Bakar berhasil ditambahkan")
		case 3:
			item := models.NewSnack("Pisang Goreng", 6000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Pisang Goreng berhasil ditambahkan")
		case 4:
			item := models.NewSnack("Cireng", 7000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Cireng berhasil ditambahkan")
			
		case 5:
			item := models.NewSnack("Seblak", 10000)
			cart.TambahkanKeranjang(item)
			fmt.Println("✅ Seblak berhasil ditambahkan")
		
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Print("Tekan ENTER untuk kembali...")
		fmt.Scanln()
		return
	}
}

