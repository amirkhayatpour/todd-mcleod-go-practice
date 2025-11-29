package main

import "fmt"

func main() {

	fmt.Println(doMath(1, 1, subtract))

}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func subtract(a int, b int) int {
	return a - b
}
