package main

import "fmt"

func sumInt(a ...int) (int, sum int) {
	for _, value := range a {
		sum += value
	}
	return len(a), sum
}

func main() {
	fmt.Println(sumInt(4, 5, 6, 7, 8))
}
