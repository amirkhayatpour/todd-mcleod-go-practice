package main

import "fmt"

func main() {
	i := 20
	for i > 0 {
		fmt.Println(i)
		i--
		if i == 10 {
			break
		}
	}
}
