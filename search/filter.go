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
		fmt.Println("🔎 Cari Menu dengan Filter\n")
		fmt.Println("Pilih kategori:")
		fmt.Println("1. Semua")
		fmt.Println("2. Makanan")
		fmt.Println("3. Minuman")
		fmt.Println("4. Snack")
		fmt.Println("0. Kembali")

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
		fmt.Println("\nUrutkan berdasarkan:")
		fmt.Println("1. Harga Tertinggi")
		fmt.Println("2. Harga Terendah")
		fmt.Println("3. Nama (A-Z)")
		fmt.Println("4. Nama (Z-A)")
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

		utils.ClearTerminal()
		fmt.Println("📋 Hasil Filter:\n")
		for i, item := range items {
			fmt.Printf("%d. %-20s Rp. %d\n", i+1, item.GetName(), item.GetHarga())
		}

		fmt.Print("\nPilih nomor untuk ditambahkan ke keranjang (0 batal): ")
		var pilih int
		fmt.Scanln(&pilih)

		if pilih > 0 && pilih <= len(items) {
			cart.TambahkanKeranjang(items[pilih-1])
			fmt.Println("✅ Berhasil ditambahkan ke keranjang.")
		} else {
			fmt.Println("❌ Dibatalkan.")
		}
		fmt.Print("Tekan ENTER untuk kembali...")
		reader.ReadString('\n')
	}
}
