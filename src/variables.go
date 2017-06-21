package main

import (
	"fmt"
	"reflect"
	"os"
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
	name := os.Getenv("USERNAME")
	course := "Learning Go"
	module := 4.4
	// pointer value of module means memory address of module
	// & => References a pointer value (to get memory address)
	// * => Dereferences a pointer value (to get the value)
	ptr := &module

	fmt.Println("Name is set to: ", name, " and is of type: ", reflect.TypeOf(name))
	fmt.Println("Course is set to: ", course, " and is of type: ", reflect.TypeOf(course))
	fmt.Println("Module is set to: ", module, " and is of type: ", reflect.TypeOf(module))
	fmt.Println("Memory address of *module* is: ", ptr, " and the value is: ", *ptr)

}
