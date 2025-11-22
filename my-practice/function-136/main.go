package main

// unfurling slice

import "fmt"

func main() {
	favoriteHeroes := []string{"Shadow fiend", "queen of pain", "gyrocoupter", "morphling"}
	echoSliceOfString(favoriteHeroes...)
}

func echoSliceOfString(text ...string) {
	for _, value := range text {
		fmt.Println(value)
	}
}
