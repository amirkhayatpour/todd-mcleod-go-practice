package main

import "fmt"

func main() {
	x()
}

var x = func() {
	fmt.Println("Shadow fiend")
}
