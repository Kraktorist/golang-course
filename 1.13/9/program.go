package main

import "fmt"

func main() {
	var (
		N, number, min, counter int
	)
	fmt.Scan(&N)
	fmt.Scan(&min)
	counter = 1
	for i := 1; i < N; i++ {
		fmt.Scan(&number)
		if number == min {
			counter++
		} else if number < min {
			counter = 1
			min = number
		}
	}
	fmt.Println(counter)
}
