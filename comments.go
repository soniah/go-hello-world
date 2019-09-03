package main

import "fmt"

func main() {
	// comment

	/* a longer,
	longer comment */

	var story string                // declare
	story = "I Like Traffic Lights" // instantiate

	story2 := "the cat sat on the mat" // string
	daysOfXmas := 12                   // integer
	const daysOfXmas2 = 12             // constant
	priceOfEggs := 3.67                // floating point

	if len(story) < 100 { // || 'or'   && 'and'   == 'equal'
		fmt.Println(story, " -> is short")
	} else if priceOfEggs > 4.0 || priceOfEggs < 0.0 {
		fmt.Println(priceOfEggs, " -> is expensive")
	} else {
		fmt.Println("what was the question?")
	}
	fmt.Println(story2, daysOfXmas) // FIN
}
