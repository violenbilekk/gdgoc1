package main

import "fmt"

func main() {
	var uang int
	var sisa int

	fmt.Scan(&uang)

	for uang != 0 {
		sisa = sisa + uang
		fmt.Scan(&uang)
	}

	fmt.Println(sisa)
}
