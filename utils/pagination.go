package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"weeklytask/models"
)

func TampilkanDenganPagination(judul string, daftar []models.Item, perPage int) models.Item {
	total := len(daftar)
	page := 0

	reader := bufio.NewReader(os.Stdin)

	for {
		ClearTerminal()
		fmt.Println()
		fmt.Println(judul)
		fmt.Println(strings.Repeat("-", 30))

		start := page * perPage
		end := start + perPage
		if end > total {
			end = total
		}

		for i := start; i < end; i++ {
			fmt.Printf("%d. %s - Rp %d\n", i+1, daftar[i].GetName(), daftar[i].GetHarga())
		}

		fmt.Println("\n> Next | < Prev | 0 Kembali")
		fmt.Print("Pilih nomor: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == ">" {
			if end < total {
				page++
			}
		} else if input == "<" {
			if page > 0 {
				page--
			}
		} else if input == "0" {
			return nil
		} else {
			num, err := strconv.Atoi(input)
			if err == nil && num > 0 && num <= total {
				return daftar[num-1]
			} else {
				fmt.Println("❌ Pilihan tidak valid")
			}
		}
	}
}
