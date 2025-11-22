package main

import "fmt"

func main() {
	justPrintSomething()
	printWhatIWant("this is just for no return and one argument function")
	myText := returnWhatIGiveYou("here is my return text")
	fmt.Println(myText)
	text, age := multiReturnFunction("here is my multi function text", 22)
	fmt.Printf("%v -------- %v", text, age)
}

func justPrintSomething() {
	fmt.Println("just no return. no param")
}

func printWhatIWant(text string) {
	fmt.Println(text)
}

func returnWhatIGiveYou(text string) string {
	return fmt.Sprint(text, "this is returned text\n")
}

func multiReturnFunction(text string, age int) (string, int) {
	return fmt.Sprint(text, "from multi return function"), age
}

/*
all available function signature
*/
