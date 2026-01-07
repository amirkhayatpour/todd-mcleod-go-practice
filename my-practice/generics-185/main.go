package main

import "fmt"

// type alias

type myType int

type myNumbers interface {
	~int | ~float64
}

func main() {
	var age myType = 31
	fmt.Println(add(age, 2.1)) // this is an error
}

func add[T myNumbers](a, b T) T {
	return a + b
}
