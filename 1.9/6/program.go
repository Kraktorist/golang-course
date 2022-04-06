package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	digit1, digit2, digit3 := a%10, a/10%10, a/100
	if digit1 == digit2 || digit2 == digit3 || digit1 == digit3 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
