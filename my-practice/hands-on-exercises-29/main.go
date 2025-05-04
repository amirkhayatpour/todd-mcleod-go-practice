package main

import "fmt"

func main() {
	for a := 0; a < 5; a++ {
		fmt.Printf("A:%v|", a)
		for b := 0; b < 5; b++ {
			fmt.Printf("B:%v", b)
		}
	}
}
