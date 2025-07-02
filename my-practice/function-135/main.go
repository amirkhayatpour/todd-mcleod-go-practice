package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func sum(numbers ...int) {
	result := 0

	for _, number := range numbers {
		result += number
	}
	fmt.Println(result)
}
