package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(mySlice, "\n---")

	for _, val := range mySlice {
		fmt.Println("RangeClause: ", val)
	}

	newSlice := []int{10, 20, 30, 40}

	// NOTE:
	// Passing arguments to ... parameters
	// Remember, since append is a variadic function, it allows us to
	// pass multiple parameters to "elems" which gets appended to it.
	// elems is a slice of type []T (in this case type []int).
	//
	// So if we instead pass a slice to it, we need to make sure that
	// we are passing on its elements. We do this by "exploding"
	// newSlice into arguments using the ... notation at the call site.
	// So, newSlice...
	//
	// https://blog.golang.org/slices
	// https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

	myName := []string{"Jewel", "Mahanta"}
	printChars(myName...) // myName is exploded to its arguments using ...
	printChars("Lightning", "McQueen")
}

func printChars(nameStr ...string) {
	for _, nameChar := range nameStr {
		fmt.Println(nameChar)
	}
}
