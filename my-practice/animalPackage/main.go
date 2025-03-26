package main

import (
	"fmt"
	"github.com/amirkhayatpour/puppy"
)

func main() {

	sound1 := puppy.Bark()
	sound2 := puppy.Barks()

	fmt.Println(sound1)
	fmt.Println(sound2)
}
