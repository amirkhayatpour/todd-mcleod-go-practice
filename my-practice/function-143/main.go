package main

import (
	"bytes"
	"fmt"
)

// byte buffer

func main() {

	b := bytes.NewBufferString("hello")
	fmt.Println(b.String())
	b.WriteString("goodbye!")
	fmt.Println(b.String())
	b.Reset()
	b.Write([]byte("i dont wanna see this is us anymore (maybe not)"))
	fmt.Println(b.String())
}
