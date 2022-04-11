package main

import "fmt"

func main() {
	var (
		n   int
		str string
	)
	fmt.Scan(&n)
	switch {
	case n%10 == 1 && n != 11:
		str = "korova"
	case 2 <= n%10 && n%10 <= 4 && n != 12 && n != 13 && n != 14:
		str = "korovy"
	default:
		str = "korov"
	}
	fmt.Printf("%v %v", n, str)
}
