package main

import "fmt"

func main() {
	fmt.Println(car("G-Class"))
}

func car(favorite string) string {
	return fmt.Sprint("this is my favorite car ", favorite)
}
