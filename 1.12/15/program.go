package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	slice := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&slice[i])
	}
	for i := 0; i < N; i = i + 2 {
		fmt.Print(slice[i], " ")
	}
}
