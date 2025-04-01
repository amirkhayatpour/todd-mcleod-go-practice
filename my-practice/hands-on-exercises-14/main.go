package main

import "fmt"

var myAge int = 30

const JOB string = "programming"

func main() {
	location := "tehran"

	fmt.Printf("my current location is: %v. (this variable is block level and declear with short variable decleartion with type %T", location, location)
	fmt.Printf("this is my age: %v. and this variable is in package scope. now i can use it on another functions. lets go", myAge)
	info()
	fmt.Printf("and this is a const variblae named location with value: %v and the type of %T. this value cant change. cuz it initialized in compile time \n", location, location)
}

func info() {
	fmt.Printf("hehe. i'm here. now i print that out from the info function. my type is %T\n", myAge)
}
