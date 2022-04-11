package main

import "fmt"

func main() {
	var (
		A, i   int
		b0, b1 int
	)
	fmt.Scan(&A)
	b1 = 1
	for i = 1; b1 <= A; i++ {
		b0, b1 = b1, b1+b0
	}
	if b0 == A {
		fmt.Println(i - 1)
	} else {
		fmt.Println(-1)
	}
}
