package main

import (
	"errors"
	"fmt"
)

func square(n int) int {
	return n * n
}

func square2(n int) (int, error) {
	if n == 42 {
		return 0, errors.New("bad number")
	}
	return n * n, nil
}

func main() {
	for i := 5; i <= 10; i++ {
		fmt.Println(i, square(i))
	}
	for i := 40; i <= 50; i++ {
		if sq, err := square2(i); err == nil {
			fmt.Println(i, sq)
		} else {
			fmt.Println("don't ask")
		}
	}
}
