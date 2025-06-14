// Launch a goroutine that processes jobs from a channel.
// Main goroutine sends jobs every 500ms.
// After 3 seconds, main sends a "quit" signal using a done channel.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, jobs chan string) {
	defer wg.Done() // Notify that this goroutine is done

	for job := range jobs {
		fmt.Println("Received job:", job)
	}
}

func master(wg *sync.WaitGroup, jobs chan string, done chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)
	timeout := time.After(3 * time.Second)

	defer ticker.Stop()
	defer wg.Done() // Notify that this goroutine is done

	for {
		select {
		case <-ticker.C:
			jobs <- "job"
		case <-timeout:
			fmt.Println("No more jobs, sending quit signal")
			close(done)
			close(jobs)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // We will have two goroutines: worker and master

	defer wg.Wait() // Wait for them to finish

	// Create channels for jobs and done signal
	done := make(chan bool)
	jobs := make(chan string)

	go worker(&wg, jobs)
	go master(&wg, jobs, done)
}
