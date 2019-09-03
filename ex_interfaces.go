package main

import (
	"fmt"
	"math"
)

type TwoDSpacer interface {
	Area() float64
	Circumference() float64
}

func Describe(t TwoDSpacer) {
	fmt.Println("Area:", t.Area())
	fmt.Println("Circumference:", t.Circumference())
}

func main() {
	Describe(Circle{2})
	Describe(Square{3})
} // FIN

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

type Square struct {
	width float64
}

func (s Square) Area() float64 {
	return s.width * s.width
}

func (s Square) Circumference() float64 {
	return 4 * s.width
}
