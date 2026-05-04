package main

import "fmt"

func main() {
	var x, y int
	var beratIkan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalSeluruhWadah float64
	var jumlahWadah int = 0

	for i := 0; i < x; i += y {
		var beratWadah float64 = 0
		for j := i; j < i+y && j < x; j++ {
			beratWadah += beratIkan[j]
		}
		fmt.Printf("%.2f ", beratWadah)
		totalSeluruhWadah += beratWadah
		jumlahWadah++
	}

	fmt.Println()

	if jumlahWadah > 0 {
		rataRata := totalSeluruhWadah / float64(jumlahWadah)
		fmt.Printf("%.2f\n", rataRata)
	}
}
