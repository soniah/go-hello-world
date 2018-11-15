package main

import "fmt"

func main() {
	const limit = 200
	balance := 100

	if balance < limit {
		shopping()
	} else {
		// TODO stop spending
	}

	for _, m := range []string{"Paris", "Rome", "New York"} {
		fmt.Println("here I am at: " + m)
	}
}

func shopping() {
	fmt.Println("spend, spend, spend")
}
