package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type employee struct {
	person
	salary int
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

	e1 := employee{
		person: father,
		salary: 50,
	}

	myFamily := []person{father, brother}

	for _, familyPerson := range myFamily {
		fmt.Printf("Name: %v\nFamily: %v\nAge: %v\n", familyPerson.first, familyPerson.last, familyPerson.age)
		fmt.Printf("%T\t%#v", familyPerson, familyPerson)
	}

	fmt.Println(e1)

	as := struct {
		model   string
		price   int
		company string
	}{
		model:   "iPhone 16e",
		price:   599,
		company: "Apple",
	}

	fmt.Printf("\n%v", as)
}
