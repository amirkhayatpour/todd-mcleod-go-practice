package main

import "fmt"

func main() {
	am := map[string]int{
		"amir":   30,
		"ali":    50,
		"masood": 90,
	}

	fmt.Println("Masood age is ", am["masood"])

	fmt.Printf("%#v\n", am)

	fmt.Println(len(am))

	lastName := make(map[string]string)

	lastName["masood"] = "shatayi"
	lastName["amir"] = "khayatpour"
	lastName["ali"] = "heydariyan"

	fmt.Println(lastName)
}
