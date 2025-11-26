package main

import "fmt"

func main() {
	fmt.Println(sumUsingVariadic(1, 2, 3, 4, 5, 6, 7, 8, 9))

	numbersSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sumUsingSlice(numbersSlice))
}

func sumUsingVariadic(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func sumUsingSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
