package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r circle) area() float64 {
	return math.Pi * math.Pow(r.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	myRectangle := rectangle{
		length: 2,
		width:  3,
	}

	myCircle := circle{
		radius: 2,
	}

	fmt.Println(info(myRectangle))
	fmt.Println(info(myCircle))
}
