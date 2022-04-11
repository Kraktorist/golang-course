package main

import "fmt"

func main() {
	var (
		N, digit, result int
	)
	fmt.Scan(&N, &digit)
	index := 1
	for ; N > 0; N /= 10 {
		if N%10 != digit {
			result = result + index*(N%10)
			index *= 10
		}
	}
	fmt.Println(result)
}
