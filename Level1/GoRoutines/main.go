// Launch 3 goroutines that print "Hello" with their ID.
// Make sure the main function waits until all goroutines finish.

package main

import (
	"fmt"
	"sync"
)

func goRoutineOne(wg *sync.WaitGroup) {
	defer wg.Done() // Notify that this goroutine is done
	fmt.Println("Hello from Go Routine One!")
}

func goRoutineTwo(wg *sync.WaitGroup) {
	defer wg.Done() // Notify that this goroutine is done
	fmt.Println("Hello from Go Routine Two!")
}

func goRoutineThree(wg *sync.WaitGroup) {
	defer wg.Done() // Notify that this goroutine is done
	fmt.Println("Hello from Go Routine Three!")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go goRoutineOne(&wg)
	go goRoutineTwo(&wg)
	go goRoutineThree(&wg)

	wg.Wait()
	fmt.Println("All Go Routines have completed execution.")
}
