package main

import "fmt"

func main() {
	var (
		x float64
	)
	fmt.Scan(&x)
	switch {
	case x > 10000:
		fmt.Printf("%e", x)
	case x <= 0:
		fmt.Printf("число %2.2f не подходит", x)
	default:
		fmt.Printf("%.4f", x*x)
	}

}
