package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favIC     []string
}

func main() {
	p1 := person{
		firstName: "Amir",
		lastName:  "Khayatpour",
		favIC:     []string{"prima", "ice pack", "strawberry"},
	}

	p2 := person{
		firstName: "Kami",
		lastName:  "chaji",
		favIC:     []string{"xar", "xak", "risheye xoshk"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for lastName, guy := range m {
		fmt.Printf("last name: %v\n%v\n", lastName, guy)
	}

}
