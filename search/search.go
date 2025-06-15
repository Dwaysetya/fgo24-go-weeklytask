package search

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"weeklytask/cart"
	"weeklytask/models"
	"weeklytask/utils"
)

func SearchMenu() {
	utils.ClearTerminal()
	fmt.Println("🔎 Cari Menu")

	fmt.Print("Masukkan keyword: ")
	reader := bufio.NewReader(os.Stdin)
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(strings.ToLower(keyword))

	var semuaItem []models.Item
	semuaItem = append(semuaItem, models.GetAllMakanan()...)
	semuaItem = append(semuaItem, models.GetAllMinuman()...)
	semuaItem = append(semuaItem, models.GetAllSnack()...)

	var hasil []models.Item
	for _, item := range semuaItem {
		if strings.Contains(strings.ToLower(item.GetName()), keyword) {
			hasil = append(hasil, item)
		}
	}

	if len(hasil) == 0 {
		fmt.Println("❌ Tidak ditemukan")
		return
	}

	fmt.Println("\nHasil pencarian:")
	for i, item := range hasil {
		fmt.Printf("%d. %s - Rp%d\n", i+1, item.GetName(), item.GetHarga())
	}

	var pilihan int
	fmt.Print("Pilih nomor untuk ditambahkan ke keranjang (0 batal): ")
	fmt.Scanln(&pilihan)

	if pilihan > 0  {
		cart.TambahkanKeranjang(hasil[pilihan-1])
		fmt.Println("✅ Berhasil ditambahkan ke keranjang.")
	} else {
		fmt.Println("❌ Dibatalkan")
	}
	fmt.Print("Tekan ENTER untuk kembali ...")
		fmt.Scanln()
}