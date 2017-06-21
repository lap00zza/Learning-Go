package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 12.234
	b := 3

	fmt.Println("type of a: ", reflect.TypeOf(a), " | type of b: ", reflect.TypeOf(b))

	// Should result in invalid operation: a + b (mismatched types float64 and int)
	// c := a + b
	// So, lets convert the type of a to an int
	c := int(a) + b
	fmt.Println("value of c: ", c, " | type of c: ", reflect.TypeOf(c))

}
