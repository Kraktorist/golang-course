package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a%400 == 0:
		fmt.Println("YES")
	case a%100 == 0:
		fmt.Println("NO")
	case a%4 == 0:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
