package main

import "fmt"
import "math/rand"

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("I ======> %v \n", i)
		switch x {
		case 0:
			fmt.Printf("x: %v \n", x)
		case 1:
			fmt.Printf("x: %v \n", x)
		case 2:
			fmt.Printf("x: %v \n", x)
		case 3:
			fmt.Printf("x: %v \n", x)
		case 4:
			fmt.Printf("x: %v \n", x)
		default:
			fmt.Printf("x is something bigger like: %v \n", x)
		}
	}

}
