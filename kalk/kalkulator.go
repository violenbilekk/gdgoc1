package main

import "fmt"

func main() {
	var i1, i2, jumlah int
	fmt.Scanln(&i1)
	jumlah = 0

	for i1 > jumlah {
		fmt.Scan(&i2)
		jumlah = jumlah + i2
		fmt.Print(i2, " ")
	}
	fmt.Println()
	fmt.Println(jumlah)
}
