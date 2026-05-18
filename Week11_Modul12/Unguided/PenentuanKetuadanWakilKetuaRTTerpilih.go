package main

import (
	"fmt"
)

func main() {
	var suara, suaraMasuk, suaraSah int
	var hitungSuara [21]int

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		suaraMasuk++

		if suara >= 1 && suara <= 20 {
			suaraSah++
			hitungSuara[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	ketua := -1
	wakil := -1
	maks1 := -1
	maks2 := -1

	// Mencari peraih suara terbanyak pertama (Ketua) dan kedua (Wakil)
	// Aturan: Jika suara sama, pilih nomor peserta yang lebih kecil (diakomodasi oleh '<' pada iterasi ascending)
	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > maks1 {
			// Geser juara 1 lama menjadi juara 2
			maks2 = maks1
			wakil = ketua

			// Perbarui juara 1 baru
			maks1 = hitungSuara[i]
			ketua = i
		} else if hitungSuara[i] > maks2 {
			// Perbarui juara 2 baru
			maks2 = hitungSuara[i]
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
