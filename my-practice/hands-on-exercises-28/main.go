package main

import "fmt"

func main() {

	for x := 0; true; x++ {

		if x%2 == 0 {
			fmt.Println(x)
		} else {
			continue
		}
	}

}
