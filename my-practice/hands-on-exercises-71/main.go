package main

import "fmt"

func main() {
	printSpecifications(myName, "Amir")
}

func printSpecifications(f func(string) string, specification string) {
	specificationToPrint := f(specification)
	fmt.Printf("this is what you want to print %s\n", specificationToPrint)
}

func myName(name string) string {
	return fmt.Sprintf("my name is: %s", name)
}
