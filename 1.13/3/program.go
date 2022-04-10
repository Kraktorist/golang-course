package main

import "fmt"

func main() {
	var (
		a int
	)
	fmt.Scan(&a)
	a0 := a % 10
	a1 := a / 10 % 10
	a2 := a / 100
	fmt.Println(a0*100 + a1*10 + a2)
}
