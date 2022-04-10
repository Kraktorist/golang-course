package main

import "fmt"

func main() {
	var (
		k int
	)
	fmt.Scan(&k)
	fmt.Printf("It is %v hours %v minutes.", k/3600, (k-k/3600*3600)/60)
}
