package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")

	ticker := time.NewTicker(time.Duration(500) * time.Millisecond)
	timer := time.NewTimer(3 * time.Second)

DONE:
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("tick tick:", t.Second())
		case fin := <-timer.C:
			fmt.Println("finishing at:", fin.String())
			break DONE
		}
	}
	fmt.Println("Done!")
} // FIN
