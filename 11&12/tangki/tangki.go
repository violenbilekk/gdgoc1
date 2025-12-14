package main

import "fmt"

func main() {
	var x, v, tambah int
	var full bool

	fmt.Scan(&x)
	v = 0
	full = false

	for full == false {
		fmt.Scan(&tambah)
		v = v + tambah

		if v >= x {
			full = true
		} else {
			full = false
		}

		fmt.Println(full)
	}
}
