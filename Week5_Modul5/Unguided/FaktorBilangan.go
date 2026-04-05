package main

import "fmt"

func cariFaktor(n, i int) {

	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Printf("%d ", i)
	}

	cariFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	cariFaktor(n, 1)
}
