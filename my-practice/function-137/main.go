package main

// defer

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Fooooo")
}

func bar() {
	fmt.Println("Baaaar")
}
