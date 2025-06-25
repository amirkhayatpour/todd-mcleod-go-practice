package main

import "fmt"

func main() {

	ages := map[string]int{
		"amir":      30,
		"abdolreza": 61,
	}

	delete(ages, "abdolreza")
	delete(ages, "shahriyar")

	fmt.Println(ages["shahriyar"])
	fmt.Println(ages)

}
