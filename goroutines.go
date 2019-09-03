package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // contrived
			continue
		}
		
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // long operation
			fmt.Println("n:", n)
		}(i)
	}
	wg.Wait()

	fmt.Println("done!")
} // FIN
