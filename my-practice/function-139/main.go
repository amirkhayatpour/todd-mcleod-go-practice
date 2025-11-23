package main

// remember interfaces (polymorphism)

import "fmt"

type human interface {
	study()
	exercise()
}

type student struct {
	firstName string
}

type clever struct {
	student
	age int
}

func (c clever) study() {
	fmt.Println(c.firstName, "is reading a book")
}

func (c clever) exercise() {
	fmt.Println(c.firstName, "is playing ping pong with ", c.age, "years old!")
}

func whatAGoodHumanDo(h human) {
	h.study()
	h.exercise()
}

func main() {
	p1 := clever{
		student: student{
			firstName: "Amir",
		},
		age: 31,
	}
	whatAGoodHumanDo(p1)
}
