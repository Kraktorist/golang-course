package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)
	var b string = strconv.Itoa(a)
	fmt.Println(string(b[len(b)-1]))
	// or
	// fmt.Println(a % 10)
}
