package main

import (
	"fmt"
	"math"
)

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Println("Luas persegi :", luas)
	fmt.Println("Keliling persegi :", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println("Luas persegi panjang :", luas)
	fmt.Println("Keliling persegi panjang :", keliling)
}

func hitungLingkaran(r float64) {
	luas := math.Pi * r * r
	keliling := 2 * math.Pi * r

	fmt.Println("Luas lingkaran :", luas)
	fmt.Println("Keliling lingkaran :", keliling)
}

func main() {
	var pilihan int

	fmt.Println("=== PROGRAM BANGUN DATAR ===")
	fmt.Println("1. Persegi")
	fmt.Println("2. Persegi Panjang")
	fmt.Println("3. Lingkaran")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var sisi int
		fmt.Print("Masukkan sisi: ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)

	case 2:
		var p, l int
		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&p)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&l)
		hitungPersegiPanjang(p, l)

	case 3:
		var r float64
		fmt.Print("Masukkan jari-jari: ")
		fmt.Scan(&r)
		hitungLingkaran(r)

	default:
		fmt.Println("Pilihan tidak valid")
	}
}
