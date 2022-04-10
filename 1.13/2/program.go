package main

import "fmt"

var (
	a int
)

func main() {
	fmt.Scan(&a)
	fmt.Println(a/100 + a/10%10 + a%10)
}
