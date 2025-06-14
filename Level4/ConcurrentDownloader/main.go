// You have a list of 10 "URLs" (just strings).
// Start 3 workers to simulate "downloading" them (sleep random 100–500ms).
// Collect and print results (URL + simulated download time).

package main

import (
	"math/rand"
	"time"
)

func downloadUrl(jobs <-chan string, result chan string) {
	for url := range jobs {
		delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
		time.Sleep(delay)
		result <- "Downloaded: " + url + " in " + delay.String()
	}
}

func main() {
	urls := []string{
		"http://example.com/1",
		"http://example.com/2",
		"http://example.com/3",
		"http://example.com/4",
		"http://example.com/5",
		"http://example.com/6",
		"http://example.com/7",
		"http://example.com/8",
		"http://example.com/9",
		"http://example.com/10",
	}

	jobs := make(chan string, len(urls))
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	result := make(chan string, len(urls))
	for i := 0; i < 3; i++ {
		go downloadUrl(jobs, result)
	}

	completed := 0
	for res := range result {
		println(res)
		completed++
		if completed == len(urls) {
			close(result)
			break
		}
	}
}
