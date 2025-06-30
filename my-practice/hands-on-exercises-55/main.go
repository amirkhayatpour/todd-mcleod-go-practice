package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int8
	color string
}

func main() {

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "XCar",
		doors: 4,
		color: "Red",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Benz",
		model: "G-class",
		doors: 4,
		color: "Black",
	}

	fmt.Printf("%v\nFuck Gas: %v\t%v\t%v\n", v1.model, v1.engine.electric, v1.doors, v1.color)
	fmt.Printf("%v\nFuck Gas: %v\t%v\t%v\n", v2.model, v2.engine.electric, v2.doors, v2.color)
}
