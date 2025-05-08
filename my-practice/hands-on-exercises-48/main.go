package main

import "fmt"

func main() {
	front := []string{"amin", "alireza", "mohsen"}
	back := []string{"alireza", "amir", "muhammad"}
	java := []string{"M.J.", "ahmad"}

	excoino := [][]string{front, back, java}

	for _, value := range excoino {
		fmt.Println(value)
		for _, members := range value {
			fmt.Println(members)
		}
	}
}
