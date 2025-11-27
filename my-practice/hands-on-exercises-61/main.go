package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("my name is %v and my age is %v", p.first, p.age)
}

func main() {
	me := person{
		first: "Amir",
		age:   31,
	}

	me.speak()

}
