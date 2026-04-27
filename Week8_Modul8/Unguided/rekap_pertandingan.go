package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var hasil []string
	i := 1

	for {
		var a, b int
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			break
		}

		if a > b {
			hasil = append(hasil, klubA)
		} else if b > a {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		i++
	}

	fmt.Println("\nHasil pertandingan:")
	for i, v := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, v)
	}
}
