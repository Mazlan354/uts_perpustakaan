package view

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"uts_perpustakaan/model"
	"uts_perpustakaan/node"
)

func ManagementBuku(reader *bufio.Reader, role string) {
	fmt.Println("== Management Buku ==")
	for {
		fmt.Println("\nMenu:")
		if role == "admin" {
			fmt.Println("1. Tambah Buku")
		}
		fmt.Println("2. Kembali")
		fmt.Print("Pilih menu: ")
		choice := readLine(reader)

		if choice == "2" {
			return
		} else if choice == "1" && role == "admin" {
			InsertBuku(reader)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func InsertBuku(reader *bufio.Reader) {
	fmt.Print("Masukkan ID Buku: ")
	id, _ := strconv.Atoi(readLine(reader))
	fmt.Print("Masukkan Judul Buku: ")
	judul := readLine(reader)
	fmt.Print("Masukkan Nama Pengarang: ")
	pengarang := readLine(reader)
	fmt.Print("Masukkan Tahun Terbit: ")
	tahun, _ := strconv.Atoi(readLine(reader))
	fmt.Print("Jumlah Halaman: ")
	halaman, _ := strconv.Atoi(readLine(reader))

	buku := node.Buku{
		ID:        id,
		Judul:     judul,
		Pengarang: pengarang,
		Tahun:     tahun,
		Halaman:   halaman,
	}

	fmt.Print("Apakah buku ini berlisensi? (ya/tidak): ")
	isLicensed := strings.ToLower(readLine(reader))
	if isLicensed == "ya" {
		model.TambahBukuBerlisensi(buku)
		fmt.Println("Buku Berlisensi berhasil ditambahkan.")
	} else {
		fmt.Print("Masukkan ID Referensi Buku Berlisensi: ")
		idRef, _ := strconv.Atoi(readLine(reader))
		if model.TambahBukuTidakBerlisensi(buku.ID, idRef) {
			fmt.Println("Buku Tidak Berlisensi berhasil ditambahkan.")
		} else {
			fmt.Println("Gagal: ID referensi tidak ditemukan.")
		}
	}
}

func ViewBukuTidakBerlisensi() {
	fmt.Println("=== Daftar Buku Tidak Berlisensi ===")
	temp := model.GetBukuTidakBerlisensi()
	for temp != nil {
		ref := model.CariBukuBerlisensi(temp.IDReferensi)
		if ref != nil {
			fmt.Printf("ID: %d | Meniru: \"%s\" oleh %s\n", temp.IDEntry, ref.Judul, ref.Pengarang)
		} else {
			fmt.Printf("ID: %d | Referensi tidak ditemukan\n", temp.IDEntry)
		}
		temp = temp.Next
	}
}

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
