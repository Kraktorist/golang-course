package main

import "fmt"

/*
func main() {
	var (
		x, p, y, year int
	)
	fmt.Scan(&x, &p, &y)
	year = 0
	for {
		year++
		x = x + x*p/100
		if x >= y {
			fmt.Println(year)
			break
		}
	}
}
*/

func main() {
	var (
		x, p, y, count int
	)
	fmt.Scan(&x, &p, &y)
	for ; x < y; x += x * p / 100 {
		count++
	}
	fmt.Println(count)
}
