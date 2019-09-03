package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Arya"
	fmt.Println("Hello", name)
	if name == "Arya" || name == "Samwell" {
		fmt.Println("Hello again")
	}
	if strings.HasPrefix(name, "A") || strings.HasPrefix(name, "S") {
		fmt.Println("Hi A/S")
	}
}
