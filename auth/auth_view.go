package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"weeklytask/utils"
)

func PromptAuth(service AuthService) {
	utils.ClearTerminal()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(`
=================================
      🔐 LOGIN TERLEBIH DAHULU      
=================================
1. 🔑 Login
2. 📝 Register
0. ❌ Keluar
=================================
`)
		fmt.Print("Pilih opsi: ")
		var opsi int
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			utils.ClearTerminal()
			fmt.Println("🔑 LOGIN")
			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')

			user := service.Login(strings.TrimSpace(email), strings.TrimSpace(password))
			if user != nil {
				fmt.Println("✅ Selamat berbelanja", user.Nama)
				return
			}
			fmt.Println("❌ Email atau password anda salah")
			fmt.Print("Tekan ENTER untuk kembali...")
			fmt.Scanln()
			utils.ClearTerminal()

		case 2:
			utils.ClearTerminal()
			fmt.Println("📝 REGISTER")
			fmt.Print("Nama: ")
			nama, _ := reader.ReadString('\n')
			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')

			service.Register(
				strings.TrimSpace(nama),
				strings.TrimSpace(email),
				strings.TrimSpace(password),
			)
			fmt.Println("✅ Registrasi berhasil")
			fmt.Print("Tekan ENTER untuk kembali...")
			fmt.Scanln()
			utils.ClearTerminal()

		case 0:
			utils.ClearTerminal()
			fmt.Println("👋 Anda keluar dari program.")
			os.Exit(0)

		default:
			utils.ClearTerminal()
			fmt.Println("❌ Pilihan tidak valid")
			fmt.Print("Tekan ENTER untuk kembali...")
			fmt.Scanln()
		}
	}
}
