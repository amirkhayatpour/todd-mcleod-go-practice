package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// write on a file using the buffer

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	p := person{
		first: "amir",
	}

	f, err := os.Create("specification.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)

	fmt.Println(b.String())

}
