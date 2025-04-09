package main

import "fmt"

func main() {
	x := 40

	switch {
	case x == 40:
		fmt.Println("x is 40")
	case x < 40:
		fmt.Println("x is smaller than 40")
	case x > 40:
		fmt.Println("x is grater than 40")
	}

	//fallthrough
	y := 50
	switch y {
	case 50:
		fmt.Println("y is 50")
		fallthrough
	case 10:
		fmt.Println("we check cuz of fallthrough!")
		fallthrough
	default:
		fmt.Println("we run default check cuz of fallthrough again")
	}
}
