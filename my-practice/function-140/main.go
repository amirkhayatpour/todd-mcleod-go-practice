package main

import "fmt"

// Stringer interface

type book struct {
	readable
	first string
	pages int
}

func (b book) String() string {
	return fmt.Sprint("hi")
}

type readable struct {
	alreadyReaded bool
}

func main() {
	b := book{
		readable: readable{
			true,
		},
		first: "binavayan",
		pages: 200,
	}

	fmt.Println(b)
}
