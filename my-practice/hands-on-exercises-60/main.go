package main

import "fmt"

func main() {

	defer foo()
	bar()

}

func foo() {
	fmt.Println("this is foo")
}

func bar() {
	fmt.Println("this is bar")
}
