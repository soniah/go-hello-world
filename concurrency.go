package main

import (
	"fmt"
)

func producer(out chan <-int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
} // FIN