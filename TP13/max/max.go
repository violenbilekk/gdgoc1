package main

import "fmt"

func main() {
	var n, besar int
	fmt.Scan(&n)
	besar = n

	for n != 0 {
		if n > besar {
			besar = n
		}
		fmt.Scan(&n)
	}
	fmt.Print(besar)

}
