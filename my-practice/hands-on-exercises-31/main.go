package main

import "fmt"

func main() {

	m := map[string]string{
		"owner": toBinary("Amir"),
		"car":   toBinary("G-class"),
	}

	for key, value := range m {
		fmt.Printf("%v => %v \n", key, value)
	}
}

func toBinary(text string) string {
	binary := ""
	for _, letter := range text {
		binary += fmt.Sprintf("%08b", letter)
	}

	return binary
}
