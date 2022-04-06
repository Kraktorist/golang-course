package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a6 := a / 100000
	a5 := a / 10000 % 10
	a4 := a / 1000 % 10
	a3 := a / 100 % 10
	a2 := a / 10 % 10
	a1 := a % 10
	if a6+a5+a4 == a3+a2+a1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
