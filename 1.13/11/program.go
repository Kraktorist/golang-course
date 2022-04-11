package main

import "fmt"

func main() {
	var (
		a, b int
	)
	fmt.Scan(&a, &b)
	for ; b%7 != 0 && b >= a; b-- {
		fmt.Println(b, b%7 != 0, b >= a, b%7 != 0 && b >= a)
	}
	if b <= a {
		fmt.Println("NO")
	} else {
		fmt.Println(b)
	}
}
