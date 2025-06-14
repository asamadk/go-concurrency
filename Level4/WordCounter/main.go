// Given a slice of strings (sentences),
// Use 4 goroutines to count the number of words in each sentence concurrently.
// Print final result as total word count.

package main

import (
	"strings"
	"sync"
	"time"
)

func countWords(wg *sync.WaitGroup, sentenceChannel chan string, totalWordChannel chan int) {
	defer wg.Done()
	for sentence := range sentenceChannel {
		time.Sleep(1 * time.Second) // Simulate expensive work
		totalWordChannel <- workCount(sentence)
	}
}

func workCount(sentence string) int {
	return len(strings.Fields(sentence))
}

func main() {
	start := time.Now() // Start timer
	var wg sync.WaitGroup

	sentences := []string{
		"Hello world",
		"Concurrency in Go is powerful",
		"Go routines are lightweight threads",
		"Channels are used for communication",
		"Synchronization is key in concurrent programming",
		"Go provides built-in support for concurrency",
		"Error handling is important in Go",
		"Go's concurrency model is based on CSP",
		"Fan out fan in is a common pattern",
		"Using goroutines can improve performance",
	}

	sentenceChannel := make(chan string, len(sentences))
	totalWordChannel := make(chan int)

	for _, sentence := range sentences {
		sentenceChannel <- sentence
	}
	close(sentenceChannel)

	wg.Add(4) // We will use 4 goroutines
	for i := 0; i < 4; i++ {
		go countWords(&wg, sentenceChannel, totalWordChannel)
	}

	go func() {
		wg.Wait()               // Wait for all goroutines to finish
		close(totalWordChannel) // Close the total word channel after all goroutines are done
	}()

	totalCount := 0
	for wordCount := range totalWordChannel {
		totalCount += wordCount
	}

	// If we did this without concurrent processing, it would take longer (10 seconds approx).
	// Here, we simulate the time taken by the goroutines to process each sentence and it takes 3 seconds approx.
	elapsed := time.Since(start)
	println("Total word count:", totalCount)
	println("Time taken:", elapsed.String())
}
