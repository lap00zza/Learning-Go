package main

import "fmt"

func main() {
	// This is how a for statement looks
	// ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .

	// This loop uses Condition. In this case it is a infinite loop
	// with the condition set to True
	/*for {
		fmt.Println("Infinite!")
	}*/

	// This loop uses ForClause
	// This loop contains a simple pre statement, a boolean expression
	// and a simple post statement
	// ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
	fmt.Println("Boolean Expression")
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	// This loop uses RangeClause
	fmt.Println("\nRange")
	nums := [4]int{1, 2, 3, 4}
	for _, val := range nums {
		fmt.Println(val)
	}
}
