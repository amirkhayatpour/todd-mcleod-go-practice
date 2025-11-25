package main

// anonymous function

import "fmt"

func main() {

	name, age := func(s string, age int) (string, int) {
		return s, age
	}("amir", 31)

	fmt.Printf("name: %v age: %v", name, age)
}
