package main

// wrapper

import (
	"fmt"
	"log"
)

type book struct {
	name string
}

func logInfo(s fmt.Stringer) {
	log.Println("this is a log from wrapper", s.String())
}

func main() {

	b1 := book{
		name: "little black fish",
	}

	logInfo(b1)
}

func (b book) String() string {
	return fmt.Sprint(b.name)
}
