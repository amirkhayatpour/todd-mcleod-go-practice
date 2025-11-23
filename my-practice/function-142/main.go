package main

import (
	"fmt"
	"log"
	"os"
)

// write files

func main() {
	f, err := os.Create("This is us best character")
	if err != nil {
		log.Printf("this is error message %s", err)
	}

	defer f.Close()

	s := []byte("Pilgrim Rick (s1 episode 8")

	characterNumber, err := f.Write(s)

	if err != nil {
		log.Printf("%s", err)
	}

	fmt.Println(characterNumber)
}
