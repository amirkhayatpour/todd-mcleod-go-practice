package main

import "fmt"

func main() {
	masood := [5]int{}

	for i := 0; i < 5; i++ {
		masood[i] = i
	}

	for _, value := range masood {
		fmt.Printf("value %v of type %T \n", value, value)
	}
}
