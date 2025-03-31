package main

import (
	"fmt"
	"github.com/amirkhayatpour/puppy"
)

func main() {

	sound1 := puppy.Bark()
	sound2 := puppy.Barks()

	sound3 := puppy.BigBark()
	sound4 := puppy.BigBarks()

	fmt.Println(sound1)
	fmt.Println(sound2)
	fmt.Println(sound3)
	fmt.Println(sound4)
}
