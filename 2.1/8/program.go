package main

import "fmt"

func vote(x int, y int, z int) int {
	return x&y | x&z | y&z
}

func main() {
	fmt.Println(vote(1, 0, 1))
}
