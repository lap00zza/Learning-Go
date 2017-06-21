package main

import (
	"fmt"
	"time"
)

func main()  {

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("BOOM!")

			// We are going to use the break statement
			// to exit the for loop as soon as we print
			// BOOM! (and not run the other code >.<)
			break
		}
		fmt.Println(timer)

		// NOTE:
		// A Duration represents the elapsed time between two instants as an
		// int64 nanosecond count.
		// To convert an integer number of units to a Duration, multiply.
		// https://golang.org/pkg/time/#Duration
		time.Sleep(500 * time.Millisecond)
	}
}
