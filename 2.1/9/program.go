package main

import "fmt"

func fibonacci(n int) int {
	phi0, phi1 := 1, 1
	for i := 2; i < n; i++ {
		phi0, phi1 = phi1, phi0+phi1
	}
	return phi1
}

func main() {
	fmt.Println(fibonacci(4))
}
