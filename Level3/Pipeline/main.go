// Stage 1: emits numbers 1-20
// Stage 2: filters even numbers
// Stage 3: squares them
// Stage 4: prints the result
// Use separate goroutines + channels for each stage.

package main

import (
	"time"
)

func generator(ch chan<- int) {
	defer close(ch) // Close the channel when done
	for i := 1; i <= 20; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func filterEven(ch chan int, evenNum chan<- int) {
	defer close(evenNum) // Close the evenNum channel when done
	for num := range ch {
		if num%2 == 0 {
			evenNum <- num
		}
	}
}

func square(evenNum chan int, squared chan<- int) {
	defer close(squared) // Close the squared channel when done
	for eve := range evenNum {
		squared <- eve * eve
	}
}

func main() {
	ch := make(chan int)
	evenNum := make(chan int)
	squared := make(chan int)

	go generator(ch)
	go filterEven(ch, evenNum)
	go square(evenNum, squared)

	for sq := range squared {
		println("Squared:", sq)
	}
}
