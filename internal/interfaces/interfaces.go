package interfaces

import (
	"fmt"
	"math"
)

type geom interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	height, width float64
}

func (g circle) area() float64 {
	return g.radius * g.radius * math.Pi
}

func (g rect) area() float64 {
	return g.height * g.width
}

func (g circle) perim() float64 {
	return 2 * g.radius * math.Pi
}

func (g rect) perim() float64 {
	return 2 * (g.height + g.width)
}

func measure(g geom) {
	fmt.Println("Area: ", g.area())
	fmt.Println("Perim:", g.perim())
}

func Interfaces() {
	fmt.Println("Interfaces package output:")
	rec := rect{3, 4}
	cir := circle{3}
	measure(rec)
	measure(cir)
	fmt.Println("")
}
