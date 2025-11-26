package main

import "fmt"

func main() {
	name, age := bar("amir", 31)
	yearsOfCoding := foo()

	fmt.Printf("my name is %s %v old, i'm coding for %v years", name, age, yearsOfCoding)
}

func foo() int {
	return 7
}

func bar(name string, age int) (string, int) {
	return name, age
}
