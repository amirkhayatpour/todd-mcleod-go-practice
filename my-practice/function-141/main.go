package main

import (
	"fmt"
	"log"
	"strconv"
)

type home struct {
	room int
}

type office struct {
	employee int
}

type capacity interface {
	countCapacity()
}

func (office office) capacity() int {
	return office.employee
}

func (home home) capacity() int {
	return home.room
}

func (office office) String() string {
	return fmt.Sprint(strconv.Itoa(office.employee))
}

func (home home) String() string {
	return fmt.Sprint(strconv.Itoa(home.room))
}

func logCapacity(s fmt.Stringer) {
	log.Println("This is Capacity: ", s.String())
}

func main() {
	myHome := home{
		room: 1,
	}

	myOffice := office{
		employee: 31,
	}

	logCapacity(myHome)
	logCapacity(myOffice)
}
