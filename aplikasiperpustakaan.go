package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type biodataBuku struct {
	kode          string
	judul         string
	pengarang     string
	penerbit      string
	jumlahhalaman int
	tahunterbit   int
}

var listBuku []biodataBuku

func main() {
	input := 0
	fmt.Println("*==Aplikasi Manajemen Daftar Buku Perpustakaan==*")
	fmt.Println("                 Silahkan Pilih:")
	fmt.Println("            1. Tambah Buku")
	fmt.Println("            2. Lihat Buku")
	fmt.Println("            3. Edit Buku")
	fmt.Println("            4. Hapus Buku")
	fmt.Println("            5. Keluar")
	fmt.Println("_===============================================_")
	fmt.Print(">> Input: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		tambahBuku()
	case 2:
		lihatBuku()
	case 3:
		editBuku()
	case 4:
		hapusBuku()
	case 5:
		os.Exit(0)
	}
	main()
}

func tambahBuku() {
	var b biodataBuku
	inputUser := bufio.NewReader(os.Stdin)
	fmt.Println("+++  Tambah Buku  +++")

	fmt.Print("Masukan Kode : ")
	b.kode, _ = inputUser.ReadString('\n')
	b.kode = strings.TrimSpace(b.kode)

	fmt.Print("Masukan Judul : ")
	b.judul, _ = inputUser.ReadString('\n')
	b.judul = strings.TrimSpace(b.judul)

	fmt.Print("Masukan Pengarang : ")
	b.pengarang, _ = inputUser.ReadString('\n')
	b.pengarang = strings.TrimSpace(b.pengarang)

	fmt.Print("Masukan Penerbit : ")
	b.penerbit, _ = inputUser.ReadString('\n')
	b.penerbit = strings.TrimSpace(b.penerbit)

	fmt.Print("Masukan Jumlah Halman : ")
	fmt.Scanln(&b.jumlahhalaman)

	fmt.Print("Masukan Tahun Terbit : ")
	fmt.Scanln(&b.tahunterbit)

	listBuku = append(listBuku, b)
	fmt.Println("Buku Berhasil Ditambahkan!")

}

func lihatBuku() {
	fmt.Println("===  Lihat Buku  ===")
	for i, b := range listBuku {
		fmt.Printf("-----------------------\nKode: %s\nJudul: %s\nPengarang: %s\nPenerbit: %s\nJumlah Halaman: %d\nTahun Terbit: %d\n",
			b.kode,
			b.judul,
			b.pengarang,
			b.penerbit,
			b.jumlahhalaman,
			b.tahunterbit,
		)
		i++
	}
	fmt.Println("-----------------------")

}

func editBuku() {
	var kodex = ""
	inputUser := bufio.NewReader(os.Stdin)
	fmt.Println("+++  Edit Buku  +++")
	fmt.Print("Masukan Kode : ")
	fmt.Scanln(&kodex)
	for i, b := range listBuku {
		if b.kode == kodex {
			fmt.Println("+++  Masukan Data Baru  +++")
			fmt.Print("Masukan Judul : ")
			listBuku[i].judul, _ = inputUser.ReadString('\n')
			listBuku[i].judul = strings.TrimSpace(listBuku[i].judul)

			fmt.Print("Masukan Pengarang : ")
			listBuku[i].pengarang, _ = inputUser.ReadString('\n')
			listBuku[i].pengarang = strings.TrimSpace(listBuku[i].pengarang)

			fmt.Print("Masukan Penerbit : ")
			listBuku[i].penerbit, _ = inputUser.ReadString('\n')
			listBuku[i].penerbit = strings.TrimSpace(listBuku[i].penerbit)

			fmt.Print("Masukan Jumlah Halman : ")
			fmt.Scanln(&listBuku[i].jumlahhalaman)

			fmt.Print("Masukan Tahun Terbit : ")
			fmt.Scanln(&listBuku[i].tahunterbit)
			fmt.Println("Buku Berhasil Diperbarui")
			main()
		}
	}
	fmt.Println("Buku dari Kode Tersebut Tidak Ditemukan!")

}

func hapusBuku() {
	var kodexx = ""
	fmt.Println("---  Hapus Buku  ---")
	fmt.Print("Kode Buku: ")
	fmt.Scanln(&kodexx)
	for i, b := range listBuku {
		if b.kode == kodexx {
			listBuku = append(
				listBuku[:i],
				listBuku[i+1:]...)

		}
	}

	fmt.Println("Buku Berhasil Dihapus!")
}
