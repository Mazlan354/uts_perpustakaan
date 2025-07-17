package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"uts_perpustakaan/view"
)

var currentRole string

func login(reader *bufio.Reader) bool {
	fmt.Print("Login sebagai (admin/pengelolabuku): ")
	inputRole, _ := reader.ReadString('\n')
	role := strings.ToLower(strings.TrimSpace(inputRole))
	if role == "admin" || role == "pengelolabuku" {
		currentRole = role
		return true
	}
	fmt.Println("Role tidak valid.")
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	if !login(reader) {
		return
	}

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Management Buku")
		fmt.Println("2. Tampilkan Buku Tidak Berlisensi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		pilihan, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Input harus berupa angka.")
			continue
		}
		switch pilihan {
		case 1:
			view.ManagementBuku(reader, currentRole)
		case 2:
			view.ViewBukuTidakBerlisensi()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
