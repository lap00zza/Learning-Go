package main

import "fmt"

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
	case "windows":
		fmt.Println("\nHere are some windows courses ...")
	default:
		fmt.Println("\nSorry! We didn't find a match.")
	}
}
