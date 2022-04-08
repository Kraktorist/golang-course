package main

import "fmt"

func main() {
	var (
		workArray [10]uint16
		j, k      int
	)
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&j, &k)
		workArray[j], workArray[k] = workArray[k], workArray[j]
	}
	for _, v := range workArray {
		fmt.Print(v, " ")
	}

}
