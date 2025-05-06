package main

import "fmt"

func main() {

	masood := []int{
		42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}

	fmt.Println(masood[:5])
	fmt.Println(masood[5:])
	fmt.Println(masood[2:7])
	fmt.Println(masood[1:6])

	/*
		● [42 43 44 45 46]
		● [47 48 49 50 51]
		● [44 45 46 47 48]
		● [43 44 45 46 47]
	*/
}
