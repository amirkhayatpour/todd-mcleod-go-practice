package main

import "fmt"

func main() {

	as := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Amir",
		friends: map[string]int{
			"Kami": 25,
			"Ali":  24,
			"Nima": 28,
		},
		favDrinks: []string{
			"Beer",
			"Wine",
			"Water",
		},
	}

	fmt.Println(as)

}
