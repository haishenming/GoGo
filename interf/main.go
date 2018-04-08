package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64{
	return 2 * r.height + 2 * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("This is geometry: ", g)
	fmt.Println("This is area: ", g.area())
	fmt.Println("This is perim: ", g.perim())
}


func main() {
	r := rect{1, 2}
	c := circle{5}

	measure(r)
	measure(c)
}
