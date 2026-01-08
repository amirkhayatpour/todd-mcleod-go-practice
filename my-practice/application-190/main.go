package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {

	p1 := person{
		Name: "amir",
		Age:  31,
	}

	p2 := person{
		Name: "tree",
		Age:  42,
	}

	persons := []person{
		p1,
		p2,
	}

	data, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

}
