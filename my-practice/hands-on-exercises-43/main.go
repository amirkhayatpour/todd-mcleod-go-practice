package main

import "fmt"

func main() {

	masood := []int{
		42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}

	for _, value := range masood {
		fmt.Printf("value %v with type of %T \n", value, value)
	}
}
