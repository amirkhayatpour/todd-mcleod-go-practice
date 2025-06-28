package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	father := person{
		first: "abdolreza",
		last:  "khayatpour dezfuli",
		age:   61,
	}

	brother := person{
		first: "shahriyar",
		last:  "khayatpour dezfuli",
		age:   22,
	}

	myFamily := []person{father, brother}

	for _, familyPerson := range myFamily {
		fmt.Printf("Name: %v\nFamily: %v\nAge: %v\n", familyPerson.first, familyPerson.last, familyPerson.age)
		fmt.Printf("%T\t%#v", familyPerson, familyPerson)
	}

}
