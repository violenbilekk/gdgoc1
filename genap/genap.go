package main

import "fmt"

func main() {
	var n, genap, digit int
	fmt.Scan(&n)
	genap = 0

	for n != 0 {
		digit = n % 10
		genap = genap + 1 - (digit % 2)
		n = n / 10
	}
	fmt.Println(genap)
}
