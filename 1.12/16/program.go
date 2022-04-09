package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	slice := make([]int, N)
	count := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&slice[i])
		if slice[i] > 0 {
			count++
		}
	}
	fmt.Println(count)
}
