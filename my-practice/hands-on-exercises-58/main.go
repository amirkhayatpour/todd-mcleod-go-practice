package main

import "fmt"

func main() {

	fmt.Printf("this is %s: %v", bar(), foo())

}

func foo() int {
	return 31
}

func bar() string {
	return "my age"
}
