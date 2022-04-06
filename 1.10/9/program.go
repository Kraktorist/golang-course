package main

import "fmt"

func main() {
	var (
		x, y string
	)
	fmt.Scan(&x, &y)
	for _, _x := range x {
		for _, _y := range y {
			if _x == _y {
				fmt.Print(string(_x), " ")
			}
		}
	}
}
