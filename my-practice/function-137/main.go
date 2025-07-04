package main

import "fmt"

func main() {
	defer thisIsTheEnd()
	skyFallMovieStarts()
}

func thisIsTheEnd() {
	fmt.Println("Hold your breath and count to 10")
}

func skyFallMovieStarts() {
	fmt.Println("007 is falling in the river")
}
