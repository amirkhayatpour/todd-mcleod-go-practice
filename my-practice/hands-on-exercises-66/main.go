package main

import "fmt"

func main() {

	myMap := map[string]int{
		"age": 22,
	}
	fmt.Println(myMap["age"])
	changedMap := changer(myMap)

	fmt.Println(myMap["age"])
	fmt.Println(changedMap["age"])

}

func changer(m map[string]int) map[string]int {
	m["age"] = 33
	return m
}
