package main

import "fmt"

type person struct {
	name string
}

func (p person) speak() {
	fmt.Println("I am ", p.name)
}

func main() {
	p1 := person{
		name: "Amir",
	}

	p2 := person{
		name: "Nima",
	}

	p1.speak()
	p2.speak()
}
