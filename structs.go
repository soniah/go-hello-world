package main

import (
	"fmt"
)

type Car struct {
	Colour     string
	SalePrice  float32
	stockPrice float32
}

type Point struct {
	x, y, z int
}

func main() {
	car1 := Car{"red", 30000.00, 25000}
	car2 := Car{stockPrice: 22000.00, SalePrice: 25500.00, Colour: "blue"}
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(Point{1, 99, 2})
	// mark1
}
