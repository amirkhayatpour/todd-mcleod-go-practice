package main

import (
	"fmt"
	"time"
)

// wrapper function
func timeOfExc(fn func()) {
	now := time.Now()
	fn()
	until := time.Since(now)
	fmt.Printf("Your function time exec is %v", until)
}

func playingDefenceOfTheAncients() {
	time.Sleep(8 * time.Second)
	fmt.Println("sf win :)")
}

func main() {
	timeOfExc(playingDefenceOfTheAncients)
}
