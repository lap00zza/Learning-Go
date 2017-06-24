package main

import (
	"fmt"
	"sync"
	"time"
)

// When main exits, all our goroutines exit.
func main() {
	// Go's concurrency model is the Actor Model; also
	// known as Communicating Sequential Processes (CSP)

	// A WaitGroup waits for a collection of goroutines to
	// finish. The main goroutine calls Add to set the
	// number of goroutines to wait for. Then each of the
	// goroutines runs and calls Done when finished. At the
	// same time, Wait can be used to block until all
	// goroutines have finished.
	// https://golang.org/pkg/sync/#WaitGroup
	var waitgroup sync.WaitGroup

	// main function (which is sort of a goroutine) calls
	// this to set the number of goroutines to wait for.
	waitgroup.Add(2)

	// Adding go keyword infront of a anonymous function
	// gives use a goroutine
	go func() {
		// Defer is used to ensure that a function call is
		// performed later in a program’s execution, usually
		// for purposes of cleanup.
		// https://gobyexample.com/defer

		// In this case, waitgroup.Done() is deferred. It means
		// that **whenever** this goroutine exits, waitgroup.Done()
		// will be called.
		defer waitgroup.Done()

		time.Sleep(2 * time.Second)
		fmt.Println("Task 1")
	}()

	go func() {
		// In this case, waitgroup.Done() is deferred. It means
		// that **whenever** this goroutine exits, waitgroup.Done()
		// will be called.
		defer waitgroup.Done()

		fmt.Println("Task 2")
	}()

	// this is needed to make sure our main function waits till
	// both our goroutines are finished. Without this, the main
	// will exit without ever waiting for the goroutines to
	// finish. And when main exits, all our goroutines exit :(
	waitgroup.Wait()
}
