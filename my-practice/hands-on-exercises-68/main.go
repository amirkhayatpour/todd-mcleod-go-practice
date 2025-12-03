package main

import "fmt"

func main() {

	func(age int, name string) {
		for i := age; i > 1; i-- {
			fmt.Printf("when you was %v, your name was %s\n", i, name)
		}
	}(31, "Amir")

}
