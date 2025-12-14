package main

import "fmt"

func main() {
	var usn, pass string
	var coba int
	coba = 0

	fmt.Scan(&usn, &pass)

	for usn != "admin" || pass != "admin" {
		fmt.Scan(&usn, &pass)
		coba = coba + 1
	}

	fmt.Println(coba)
	fmt.Println("login Berhasil")
}
