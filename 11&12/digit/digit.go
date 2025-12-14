package main

import "fmt"

func main() {
	var digit, sat, jumlah int

	fmt.Scanln(&digit)
	for digit > 0 {
		sat = digit % 10
		digit = digit / 10
		fmt.Print(sat, " ")
		jumlah = jumlah + sat
	}
	fmt.Println()
	fmt.Println(jumlah)
}
