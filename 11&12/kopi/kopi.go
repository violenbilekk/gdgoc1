package main

import "fmt"

func main() {
	var n, m, x, y, cangkir int
	fmt.Scan(&n, &m, &x, &y)
	for x <= n && y <= m {
		n = n - x
		m = m - y
		cangkir++
	}
	fmt.Print(cangkir)
}
