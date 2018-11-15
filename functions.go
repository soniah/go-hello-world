package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// multiple inputs, same type; named returns
func plusPlus(a, b, c int) (result int, err error) {
	// many long calculations
	result = a + b + c
	// many long calculations
	return result, nil
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	if res2, err := plusPlus(1, 2, 3); err != nil {
		fmt.Println("1+2+3 =", res2)
	}
}
