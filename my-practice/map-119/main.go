package main

import "fmt"

func main() {

	capitals := map[string]string{
		"japan":   "tokyo",
		"iran":    "tehran",
		"england": "london",
	}

	if capital, ok := capitals["japan"]; ok {
		fmt.Println(capital)
	} else {
		fmt.Println("it's not defined!")
	}

	if italyCapital, ok := capitals["italy"]; ok {
		fmt.Println(italyCapital)
	} else {
		fmt.Println("wooops!")
	}
}
