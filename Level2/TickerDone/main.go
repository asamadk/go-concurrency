// Print "tick" every second, stop after 5 ticks using a done channel.
// Use time.NewTicker and select with a done channel.

package main

import "time"

func tick(countChan chan int, done chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			println("tick")
			countChan <- 1
		case <-done:
			println("Ticker stopped")
			return
		}
	}
}

func main() {
	countChan := make(chan int)
	done := make(chan bool)
	go tick(countChan, done)

	count := 0
	for {
		select {
		case <-countChan:
			count++
			if count >= 5 {
				done <- true
				return
			}
		}
	}
}
