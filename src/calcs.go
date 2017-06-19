package main

import (
	"fmt"
	"reflect"
)

func main ()  {
	a := 12.234
	b := 3

	fmt.Println("type of a: ", reflect.TypeOf(a), " | type of b: ", reflect.TypeOf(b))

	c := a + b
	fmt.Println("value of c: ", c, " | type of c: ", reflect.TypeOf(c))
}
