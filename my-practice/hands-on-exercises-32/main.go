package main

import (
	"fmt"
)

func main() {

	data := map[string]string{
		"owner": stringToBinary("Amir"),
		"car":   stringToBinary("G-class"),
	}

	fmt.Println(data["owner"])
	if age, ok := data["age"]; ok {
		fmt.Println(age)
	} else {
		fmt.Println("I cant find it")
	}
}

func stringToBinary(text string) string {
	binary := ""
	for _, letter := range text {
		binary += fmt.Sprintf("%08b", letter)
	}

	return binary
}
