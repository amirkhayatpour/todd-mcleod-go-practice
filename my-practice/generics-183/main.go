package main

import "fmt"

// generics (DRY)

func main() {
	fmt.Println(add(1, 1.1))
}

func add[T int | float64](a, b T) T {
	return a + b
}
