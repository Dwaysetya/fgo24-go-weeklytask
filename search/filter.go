package search

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"weeklytask/cart"
	"weeklytask/models"
	"weeklytask/utils"
)

func TampilFilterMenu() {
	utils.ClearTerminal()
	reader := bufio.NewReader(os.Stdin)

	for {
		utils.ClearTerminal()
		utils.ClearTerminal()
		fmt.Println(`
=================================
	🔎  FILTER MENU PENCARIAN    
=================================
1. 🍽️  Semua Menu
2. 🍛  Makanan
3. 🥤  Minuman
4. 🍟  Snack
0. ❌  Kembali
=================================
		`)

		fmt.Print("Masukkan pilihan: ")
		var kategori int
		fmt.Scanln(&kategori)

		if kategori == 0 {
			return
		}

		var items []models.Item
		switch kategori {
		case 1:
			items = append(items, models.GetAllMakanan()...)
			items = append(items, models.GetAllMinuman()...)
			items = append(items, models.GetAllSnack()...)
		case 2:
			items = models.GetAllMakanan()
		case 3:
			items = models.GetAllMinuman()
		case 4:
			items = models.GetAllSnack()
		default:
			fmt.Println("❌ Pilihan kategori tidak valid.")
			fmt.Print("Tekan ENTER untuk kembali...")
			reader.ReadString('\n')
			continue
		}

		utils.ClearTerminal()
		fmt.Println(`
=================================
	🧮  URUTKAN BERDASARKAN      
=================================
1. 💰 Harga Tertinggi
2. 🪙 Harga Terendah
3. 🔤 Nama (A-Z)
4. 🔡 Nama (Z-A)
0. ❌ Batal
=================================
		`)
		fmt.Print("Masukkan pilihan filter: ")
		var filter int
		fmt.Scanln(&filter)

		switch filter {
		case 1:
			sort.Slice(items, func(i, j int) bool {
				return items[i].GetHarga() > items[j].GetHarga()
			})
		case 2:
			sort.Slice(items, func(i, j int) bool {
				return items[i].GetHarga() < items[j].GetHarga()
			})
		case 3:
			sort.Slice(items, func(i, j int) bool {
				return strings.ToLower(items[i].GetName()) < strings.ToLower(items[j].GetName())
			})
		case 4:
			sort.Slice(items, func(i, j int) bool {
				return strings.ToLower(items[i].GetName()) > strings.ToLower(items[j].GetName())
			})
		default:
			fmt.Println("❌ Pilihan filter tidak valid.")
			fmt.Print("Tekan ENTER untuk kembali...")
			reader.ReadString('\n')
			continue
		}

		for {
			utils.ClearTerminal()
			item := utils.TampilkanPagination("📋 Hasil Filter", items, 5)

			if item == nil {
				break
			}

			cart.TambahkanKeranjang(item)
			fmt.Printf("✅ %s berhasil ditambahkan ke keranjang\n", item.GetName())
			fmt.Print("Tekan ENTER untuk kembali ke hasil filter...")
			reader.ReadString('\n')
		}
	}
}
