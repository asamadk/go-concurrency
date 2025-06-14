// Create a generator that emits numbers from 1 to 10 with 100ms delay.
// Consume and print them from the main function.

package main

import (
	"fmt"
	"time"
)

// Using unidirectional channels to create a generator that emits numbers
// Generator can only send values to the channel
// Consumer can only receive values from the channel
// This pattern is useful for separating concerns and making the code cleaner.
func generator(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go generator(ch)

	for c := range ch {
		fmt.Println("Received:", c)
	}
}
