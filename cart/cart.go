package cart

import (
	"fmt"
	"sync"
	"weeklytask/models"
	"weeklytask/utils"
)

var Keranjang []models.Item

func TambahkanKeranjang(item models.Item) {
	Keranjang = append(Keranjang, item)
}

func hitungTotal(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	var total int = 0
	for _, item := range Keranjang {
		total += item.GetHarga()
	}
	ch <- total
}

func LihatKeranjang() int {
	utils.ClearTerminal()

	if len(Keranjang) == 0 {
		fmt.Println("🛒 Anda belum membeli apapun")
		fmt.Print("Tekan ENTER untuk kembali...")
		fmt.Scanln()
	}

	fmt.Println("📦 Daftar Isi Keranjang:")
	for i, item := range Keranjang {
		fmt.Printf("%d. %s Rp.%d \n", i+1, item.GetName(), item.GetHarga())
	}

	var wg sync.WaitGroup
	totalChan := make(chan int, 1)

	wg.Add(1)
	go hitungTotal(&wg, totalChan)

	total := <-totalChan 
	wg.Wait()
	close(totalChan)

	fmt.Printf("\n💰 Total: Rp.%d\n", total)
	return total
}


func Checkout() {
	utils.ClearTerminal()

	if len(Keranjang) == 0 {
		fmt.Println("🛒 Keranjang masih kosong, tidak bisa checkout")
		fmt.Print("Tekan ENTER untuk kembali...")
		fmt.Scanln()
		return
	}

	total := LihatKeranjang()

	var konfirmasi string
	fmt.Print("Yakin ingin melakukan checkout? (y/n): ")
	fmt.Scanln(&konfirmasi)

	if konfirmasi != "y" && konfirmasi != "Y" {
		fmt.Println("❌ Checkout dibatalkan")
		fmt.Print("Tekan ENTER untuk kembali...")
		fmt.Scanln()
		return
	}
	utils.ClearTerminal()
	fmt.Println("===========================")
	fmt.Println("        STRUK BELANJA      ")
	fmt.Println("===========================")
	for i, item := range Keranjang {
		fmt.Printf("%d. %s - Rp.%d\n", i+1, item.GetName(), item.GetHarga())
	}
	fmt.Println("---------------------------")
	fmt.Printf("Total Bayar: Rp.%d\n", total)
	fmt.Println("===========================")
	fmt.Println("Terima kasih 🙏")
	fmt.Println()

	Keranjang = []models.Item{}

	fmt.Print("Tekan ENTER untuk kembali ke menu utama...")
	fmt.Scanln()
}

