package main

import "fmt"

func main() {
	var target, n, letak int
	var found, done bool

	fmt.Scan(&target)

	for !done {
		fmt.Scan(&n)

		if n == 0 {
			done = true
		} else {
			letak++
			if n == target {
				found = true
				done = true
			}
		}
	}

	if found {
		fmt.Println(letak)
	} else {
		fmt.Println("TIDAK ADA")
	}
}
