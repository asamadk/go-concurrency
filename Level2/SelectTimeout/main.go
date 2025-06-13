// Create a goroutine that sleeps for 3 seconds and then sends a value.
// Use a select block to wait for either the value or a 2-second timeout.

package main

import (
	"fmt"
	"time"
)

func sleepAndSend(valChan chan int) {
	fmt.Println("Goroutine started, sleeping for 3 seconds...")
	time.Sleep(3 * time.Second)

	fmt.Println("Goroutine woke up, sending value to channel...")
	valChan <- 42
}

func main() {
	valChan := make(chan int)

	go sleepAndSend(valChan)

	select {
	case value := <-valChan:
		println("Received value:", value)
	case <-time.After(5 * time.Second):
		println("Timeout occurred, no value received")
		println("The goroutine took too long to send a value.")
		println("Exiting the program due to timeout.")
	}
}
