package main

import "fmt"

// type set interface

type myNumbers interface {
	int | float64
}

func add[T myNumbers](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2.2))
}
