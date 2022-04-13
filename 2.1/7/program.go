package main

import "fmt"

func minimumFromFour() int {
	var a int
	fmt.Scan(&a)
	min := a
	for i := 1; i < 4; i++ {
		fmt.Scan(&a)
		if a < min {
			min = a
		}
	}
	return min
}

func main() {

}
