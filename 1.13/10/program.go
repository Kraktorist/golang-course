package main

import "fmt"

func sum(number int) int {
	var result int
	for ; number > 0; number /= 10 {
		result += number % 10
	}
	if result > 9 {
		result = sum(result)
	}
	return result
}

func main() {
	var (
		number int
	)
	fmt.Scan(&number)
	fmt.Println(sum(number))
}

/*
	1 подсчитать сумму цифр
	2 Если сумма цифр меньше 10, то вывести ее
	3 если сумма >=10, то пункт 1 для этой суммы
*/
