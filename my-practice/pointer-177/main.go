package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) getFirst() string {
	return p.first
}

func (p *person) getAge() int {
	return p.age
}

func main() {
	p1 := person{
		first: "Amir",
		age:   31,
	}

	p2 := &person{
		first: "koala",
		age:   25,
	}

	fmt.Println(p1.getFirst())
	fmt.Println(p2.getAge())
	fmt.Println(p1.getFirst())
	fmt.Println(p2.getAge())

}
