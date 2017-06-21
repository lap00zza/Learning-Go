package main

import "fmt"

func main() {
	bestFinish := bestLeagueFinishes(10, 12, 16, 13, 11, 8, 7, 21, 20, 6)
	fmt.Println("\nBest Finish: ", bestFinish)
}

// bestLeagueFinishes accepts any number of int's (denoted by ...int).
// The values passed onto this function is turned into a slice and is
// saved in 'finishes'.
func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]

	// NOTE:
	// range iterates over elements in a variety of data.
	// It returns 2 values for every iteration - index
	// and value. If we don't need the index, we can
	// ignore it with the blank identifier _
	// src: https://gobyexample.com/range
	for _, value := range finishes {
		if value < best {
			best = value
		}
	}

	return best
}
