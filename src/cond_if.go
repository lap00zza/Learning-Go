package main

import "fmt"

func main() {
	//firstRank := 39
	//secondRank := 100

	// go also support simple statements in if.
	// IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
	// Simple statements: https://golang.org/ref/spec#SimpleStmt
	if firstRank, secondRank := 39, 100; firstRank < secondRank {
		fmt.Println("\nFirst is doint better!")
	} else if firstRank > secondRank {
		fmt.Println("\nSecond is doing better!")
	} else {
		fmt.Println("\nSomething fishy is going on!")
	}
}
