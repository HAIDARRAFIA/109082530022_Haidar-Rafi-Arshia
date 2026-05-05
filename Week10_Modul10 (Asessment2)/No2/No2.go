package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM, nama string
	nilai     int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(T arrayMahasiswa, n int, nim string) int {
	for i := 1; i <= n; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, nim string) int {
	max := -1
	for i := 1; i <= n; i++ {
		if T[i].NIM == nim {
			if max == -1 || T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}

func main() {
	var T arrayMahasiswa
	var n int
	var searchNIM string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari : ")
	fmt.Scan(&searchNIM)

	first := cariNilaiPertama(T, n, searchNIM)
	highest := cariNilaiTerbesar(T, n, searchNIM)

	if first != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", searchNIM, first)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", searchNIM, highest)
	} else {
		fmt.Println("NIM tidak ditemukan.")
	}
}
