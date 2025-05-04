package main

import "fmt"
import "math/rand"

func main() {
	for round := 0; round < 100; round++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println(x)
		}
	}
}
