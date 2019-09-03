package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := []int{1,2,3}
	var sum int
	for _, n := range numbers {
		sum += n
	}
	cities := []string{"Sydney", "Melbourne", "Lismore"}
	var result string
	for _, city := range cities {
		result += city
	}
	fmt.Println(sum, "\n", result, "\n", strings.Join(cities, ", "))
}
