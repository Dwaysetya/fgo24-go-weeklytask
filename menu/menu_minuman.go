package menu

import (
	"fmt"
	"weeklytask/cart"
	"weeklytask/models"
	"weeklytask/utils"
)

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