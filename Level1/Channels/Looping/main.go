// Send 10 numbers into a buffered channel, close it,
// and print each value using a range loop over the channel.

package main

import "fmt"

func main() {
	bufferedChannel := make(chan int, 10)

	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3
	bufferedChannel <- 4
	bufferedChannel <- 5
	bufferedChannel <- 6
	bufferedChannel <- 7
	bufferedChannel <- 8
	bufferedChannel <- 9
	bufferedChannel <- 10

	close(bufferedChannel)

	for value := range bufferedChannel {
		fmt.Println(value)
	}
}
