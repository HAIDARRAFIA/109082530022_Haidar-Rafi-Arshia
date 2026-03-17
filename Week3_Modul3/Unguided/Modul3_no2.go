package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var x, y, z int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	fmt.Print("Masukkan nilai z: ")
	fmt.Scan(&z)

	fmt.Println("f(g(h(x))) =", f(g(h(x))))
	fmt.Println("g(h(f(y))) =", g(h(f(y))))
	fmt.Println("h(f(g(z))) =", h(f(g(z))))
}
