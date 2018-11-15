package main

import "fmt"

func main() {
	mySlice := []string{"Paris", "Rome"}
	mySlice = append(mySlice, "Rio")
	fmt.Println(mySlice)
	for i, value := range mySlice {
		fmt.Println("i", i, "value", value)
	}
}
