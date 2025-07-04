package main

import "fmt"

type benz struct {
	model string
	myFav bool
}

type toyota struct {
	model string
	myFav bool
}

type car interface {
	drive()
}

func (auto benz) drive() {
	fmt.Println(auto.model, " is moving ...... ! Am i like it ? ", auto.myFav)
}

func (auto toyota) drive() {
	fmt.Println(auto.model, " is moving ...... ! Am i like it ? ", auto.myFav)
}

func letGo(car car) {
	car.drive()
}

func main() {
	gClass := benz{
		model: "G-class",
		myFav: true,
	}

	landCruiser := toyota{
		model: "Land Cruiser",
		myFav: false,
	}

	letGo(gClass)
	letGo(landCruiser)
}
