package cart

import (
	"fmt"
	"sync"
	"time"
	"weeklytask/models"
	"weeklytask/utils"
)

var Keranjang []models.Item

func TambahkanKeranjang(item models.Item) {
	Keranjang = append(Keranjang, item)
}

func hitungTotal(wg *sync.WaitGroup, ch chan int64) {
	defer wg.Done()

	var total int64 = 0
	for _, item := range Keranjang {
		total += item.GetHarga()
	}
	ch <- total
}

func LihatKeranjang() int64 {
	utils.ClearTerminal()

	if len(Keranjang) == 0 {
		fmt.Println("🛒 Keranjang kosong.")
		return 0
	}

	fmt.Println("📦 Daftar Isi Keranjang:\n")
	for i, item := range Keranjang {
		fmt.Printf("%d. %s Rp.%d \n", i+1, item.GetName(), item.GetHarga())
	}

	var wg sync.WaitGroup
	totalChan := make(chan int64, 1) // BUFFERED supaya tidak deadlock

	wg.Add(1)
	go hitungTotal(&wg, totalChan)

	total := <-totalChan // langsung terima
	wg.Wait()
	close(totalChan)

	fmt.Printf("\n💰 Total: Rp.%d\n", total)
	return total
}


func Checkout() {
	utils.ClearTerminal()

	if len(Keranjang) == 0 {
		fmt.Println("🛒 Keranjang masih kosong, tidak bisa checkout.")
		fmt.Print("Tekan ENTER untuk kembali...")
		fmt.Scanln()
		return
	}

	total := LihatKeranjang()

	var konfirmasi string
	fmt.Print("Yakin ingin melakukan checkout? (y/n): ")
	fmt.Scanln(&konfirmasi)

	if konfirmasi != "y" && konfirmasi != "Y" {
		fmt.Println("❌ Checkout dibatalkan.")
		fmt.Print("Tekan ENTER untuk kembali...")
		fmt.Scanln()
		return
	}
	utils.ClearTerminal()
	fmt.Println("\n🧾 Sedang memproses checkout...")
	time.Sleep(2 * time.Second)
	utils.ClearTerminal()
	fmt.Printf("✅ Checkout berhasil. Total belanja: Rp.%d\n", total)
	fmt.Println("Terima kasih telah berbelanja!")

	Keranjang = []models.Item{}

	fmt.Print("Tekan ENTER untuk kembali ke menu utama...")
	fmt.Scanln()
}

