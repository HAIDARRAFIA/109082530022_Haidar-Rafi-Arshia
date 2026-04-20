package main

import "fmt"

type angka int
type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var b Buku
	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku: ")
	fmt.Scan(&b.judul)
	fmt.Print("Masukkan nama penulis: ")
	fmt.Scan(&b.penulis)
	fmt.Print("Masukkan penerbit: ")
	fmt.Scan(&b.penerbit)
	fmt.Print("Masukkan tahun terbit: ")
	fmt.Scan(&b.tahunTerbit)
	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&b.jumlahHalaman)

	fmt.Printf("\nBIODATA BUKU\nJudul Buku: %s\nPenulis: %s\nPenerbit: %s\nTahun Terbit: %d\nJumlah Halaman: %d\n",
		b.judul, b.penulis, b.penerbit, b.tahunTerbit, b.jumlahHalaman)
}
