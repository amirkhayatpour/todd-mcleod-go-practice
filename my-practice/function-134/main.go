package main

import "fmt"

func main() {
	sayHallo()
	sayMyName("Heisenberg")
	product := whoMakesWhat("Heisenberg")
	fmt.Printf("who makes what: %v", product)
	fullName, age := ageFinder("Amir", "Khayatpour")
	fmt.Printf("person: %v found age: %v", fullName, age)
}

func sayHallo() {
	fmt.Println("Hallo")
}

func sayMyName(name string) {
	fmt.Printf("%v, You god damn right! >:)", name)
}

func whoMakesWhat(person string) string {
	switch person {
	case "Heisenberg":
		return "crystal"
	default:
		return "I dont know!"
	}
}

func ageFinder(name string, last string) (string, int) {
	fullName := fmt.Sprintf("%s %s", name, last)
	var age int
	switch fullName {
	case "Amir Khayatpour":
		age = 30
	default:
		age = 0
	}

	return fullName, age
}
