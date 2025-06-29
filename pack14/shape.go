package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	var s Shape

	s = Rectangle{Width: 10, Height: 5}
	fmt.Printf("Площадь прямоугольника: %.2f\n", s.Area())

	s = Circle{Radius: 5}
	fmt.Printf("Площадь круга: %.2f\n", s.Area())
}
