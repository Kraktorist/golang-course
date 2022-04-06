package main

import "fmt"

func main() {
	var (
		a   int
		b   int
		sum int = 0
	)
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
