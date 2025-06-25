package main

import "fmt"

func main() {

	aMap := map[string]int{
		"amir":   30,
		"ali":    77,
		"masood": 100,
	}

	for key, value := range aMap {
		fmt.Println(key, value)
	}

	for _, value := range aMap {
		fmt.Println(value)
	}

	aSlice := []string{"Koala", "37"}

	for key, value := range aSlice {
		println(key, value)
	}
}
