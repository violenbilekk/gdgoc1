package main

import "fmt"

func main() {
	var x, sat1, sat2 int
	var selisih bool

	fmt.Scan(&x)
	sat1 = x % 10
	selisih = false

	for x != 0 {
		x = x - sat1
		sat2 = x % 10
		if sat1-sat2 == 1 || sat2-sat1 == 1 {
			selisih = true
		}
	}
	fmt.Print(selisih)

}
