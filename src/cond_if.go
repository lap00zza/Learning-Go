package main

import "fmt"

func main() {
	firstRank := 39
	secondRank := 100

	if firstRank < secondRank {
		fmt.Println("\nFirst is doint better!")
	} else if firstRank > secondRank {
		fmt.Println("\nSecond is doing better!")
	} else {
		fmt.Println("\nSomething fishy is going on!")
	}
}
