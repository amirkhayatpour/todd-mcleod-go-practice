package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type myNumbers interface {
	constraints.Integer | constraints.Float
}

func main() {
	fmt.Println(add(1, 1.2))
}

func add[T myNumbers](a, b T) T {
	return a + b
}
