package main

import (
	"fmt"
	"reflect"
)

// These are variables declared at the package level.
// These are available to all functions. (So scope: Global)
var (
	//name, course string
	//module       float64

	// Go an infer types from the initialization
	fullName = "Jewel Mahanta"
	score    = 4.3
)

func main() {
	// Automatic type inference
	fmt.Println("Name: ", fullName, " and is of type: ", reflect.TypeOf(fullName))
	fmt.Println("Score: ", score, " and is of type: ", reflect.TypeOf(score))

	// Variables declared in the function level can be declared
	// and initialized using the shorthand `:=`.
	name := "Jewel Mahanta"
	//course := "Learning Go"
	module := 4.4
	fmt.Println("Name is set to: ", name, " and is of type: ", reflect.TypeOf(name))
	fmt.Println("Module is set to: ", module, " and is of type: ", reflect.TypeOf(module))

}
