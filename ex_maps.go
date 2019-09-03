package main

import "fmt"

func main() {
	students := map[string]int{"Bill": 28, "Mary": 74, "Juana": 91, "John": 79}
	for name, score := range students {
		if 60 <= score && score <= 80 {
			fmt.Printf("Name: %s Score: %d\n", name, score)
		}
	}

	cities := map[string]string{"Mullum": "NSW", "Ballina": "NSW", "1770": "QLD"}
	counts := make(map[string]int)
	for _, state := range cities {
		counts[state] += 1
	}
	fmt.Println(counts)
}
