// Create one unbuffered and one buffered channel.
// Try sending to both without a receiver.
// Observe blocking behavior and write down your findings.

package main

import "fmt"

func main() {
	bufferedChannel := make(chan string, 3)
	unBufferedChannel := make(chan string)

	fmt.Println("Sending messages to channels...")
	// Sending messages to the buffered channel

	bufferedChannel <- "Buffered Channel Message 1"
	bufferedChannel <- "Buffered Channel Message 2"
	bufferedChannel <- "Buffered Channel Message 3"
	// The above lines will not block because the buffered channel can hold up to 3 messages.

	fmt.Println("Buffered channel messages sent successfully.")

	unBufferedChannel <- "Unbuffered Channel Message 1"
	// The above line will block until a receiver is ready to receive the message.

	fmt.Println("Unbuffered channel message sent successfully.")
	// To avoid deadlock, we need to close the buffered channel
}

// Observations:
// 1. The buffered channel allows sending multiple messages without blocking until the buffer is full.
// 2. The unbuffered channel blocks the sending goroutine until a receiver is ready to receive the message.
// 3. If we try to send more messages to the buffered channel than its capacity, it will block until some messages are received.
// 4. If we do not have a receiver for the unbuffered channel, the program will deadlock at that point.
// 5. Always ensure to close channels when they are no longer needed to avoid resource leaks.
// 6. Using buffered channels can help in scenarios where you want to decouple the sender and receiver, allowing for more flexibility in goroutine execution.
// 7. Unbuffered channels are useful for synchronization between goroutines, ensuring that a send operation waits for a corresponding receive operation.
// 8. It's important to handle channel operations carefully to avoid deadlocks and ensure proper synchronization in concurrent programs.
// 9. When using channels, always consider the potential for blocking and deadlocks, especially in complex applications with multiple goroutines.
