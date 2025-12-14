package main

import "fmt"

func main() {
	var input byte
	var cek int
	cek = 0
	for input != '#' {
		cek = cek + 1
		fmt.Scanf("%c", &input)
	}
	fmt.Println(cek - 1)
}
