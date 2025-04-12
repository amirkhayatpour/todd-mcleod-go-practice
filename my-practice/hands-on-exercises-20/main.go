package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(500)

	switch {
	case x <= 100:
		fmt.Println("between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("between 101 and 200")
	case x > 200 && x <= 250:
		fmt.Println("between 200 and 250")
	default:
		fmt.Println("value is grater than 250")
	}
}
