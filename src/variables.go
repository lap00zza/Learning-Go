package main

import (
	"fmt"
	"reflect"
)

var (
	name, course string
	module float64
)

func main ()  {
	fmt.Println("Name is set to: ", name, " and is of type: ", reflect.TypeOf(name))
	fmt.Println("Course is set to: ", course, " and is of type: ", reflect.TypeOf(course))
	fmt.Println("Module is set to: ", module, " and is of type: ", reflect.TypeOf(module))
}
