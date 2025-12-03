package main

import (
	"fmt"
	"math"
)

func main() {
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powinator(n float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(n, c)
	}
}
