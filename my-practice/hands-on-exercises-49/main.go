package main

import "fmt"

func main() {

	interests := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"do_dr":            {"cats", "ice cream", "sunsets"},
	}

	interests["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for key, value := range interests {
		fmt.Printf("%v\n", key)
		for _, interest := range value {
			fmt.Printf("%v\t", interest)
		}
		fmt.Printf("\n")
	}

}
