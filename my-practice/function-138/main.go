package main

// methods

import "fmt"

type person struct {
	firstName string
}

func (p person) speak() {
	fmt.Println("My name is: ", p.firstName)
}

func main() {
	p1 := person{
		firstName: "amir",
	}

	p2 := person{
		firstName: "koala",
	}

	p1.speak()
	p2.speak()
}
