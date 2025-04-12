package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := Intn(250)

	fmt.Println(x)

	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x > 101 && x < 200 {
		fmt.Println("between 100 and 200")
	} else if x > 200 {
		fmt.Println("between 200 and 250")
	}

}

func Intn(n int) int {
	return rand.Intn(n)
}
