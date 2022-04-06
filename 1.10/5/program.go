package main

import "fmt"

func main() {
	var (
		a     int
		max   int = 0
		count int = 0
	)
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > max {
			max = a
			count = 1
		} else if a == max {
			count++
		}
	}
	fmt.Println(count)
}
