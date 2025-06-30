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

	boys := []person{
		p1,
		p2,
	}

	for _, boy := range boys {
		fmt.Println(boy.firstName)
		for _, favIces := range boy.favIC {
			fmt.Printf("\n favorite ice creame: %v", favIces)
		}
		fmt.Printf("\n")
	}
}
