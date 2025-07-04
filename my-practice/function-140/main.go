package main

import (
	"fmt"
	"strconv"
)

type book struct {
	name string
}

type age int

func (b book) String() string {
	return fmt.Sprint("This is the book name: ", b.name)
}

func (a age) String() string {
	return fmt.Sprint("And this is my Age: ", strconv.Itoa(int(a)))
}

func main() {
	b := book{
		name: "The Miserable Ones",
	}

	var n age = 30

	fmt.Println(b)
	fmt.Println(n)
}
