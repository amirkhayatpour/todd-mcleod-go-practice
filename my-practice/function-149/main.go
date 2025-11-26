package main

import "fmt"

// closure

func main() {

	f := incrementor() // func() int
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
