package main

import (
	"fmt"
	"strings"
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

	fmt.Println(car1.Format(">>> ")) // mark4
	fmt.Println(car2.Format(">>> "))
	fmt.Println(Point{1, 99, 2}.Format("PT:")) // mark5
}

type Formatter interface { // mark2
	Format(prefix string) string
}

func (c Car) Format(prefix string) string {
	indent := "                         "[0:len(prefix)] // HACK
	return prefix + fmt.Sprintf("Colour: %s\n%sSalePrice: $%.2f\n%sstockprice: xxxxx",
		strings.ToUpper(c.Colour), indent, c.SalePrice, indent)
}

func (p Point) Format(prefix string) string {
	return fmt.Sprintf("%s[%d:%d:%d]", prefix, p.x, p.y, p.z)
} // mark3
