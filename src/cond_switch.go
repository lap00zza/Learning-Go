package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// NOTE:
	// Switch also supports simple statements
	// ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
	// In go, switch does not have implicit fallthrough. So once there is a match,
	// it exits the entire switch block and moves on to the next line.
	switch course := "docker"; course {
	case "linux":
		fmt.Println("\nHere are some linux courses ...")
	case "docker":
		fmt.Println("\nHere are some docker courses ...")

		// This keyword allows fallthrough behaviour.
		// So in this case, after matching case docker,
		// it will execute the docker block and then
		// execute the windows block too.
		fallthrough
	case "windows":
		fmt.Println("\nHere are some windows courses ...")
	default:
		fmt.Println("\nSorry! We didn't find a match.")

	}

	// In go, it is more idiomatic to match multiple cases like
	// shown below than using fallthrough. So the approach below
	// is always preferred.
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Even number: ", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd number: ", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
