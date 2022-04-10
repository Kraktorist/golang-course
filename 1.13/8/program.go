package main

import "fmt"

func main() {
	var (
		N, number, counter int
	)
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&number)
		if number == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
