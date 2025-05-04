package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47}
	for key, value := range x {
		fmt.Printf("key: %v | value: %v \n", key, value)
	}
}
