package main

//variadic parameter

import "fmt"

func main() {

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func sum(numbers ...int) {
	n := 0
	for index, number := range numbers {
		fmt.Printf("index: %v ---- number is : %v ---- and sum : %v\n", index, number, n)
		n += number
	}
}
