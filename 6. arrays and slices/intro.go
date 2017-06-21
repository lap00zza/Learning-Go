package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Arrays and slices are numbered lists of same datatype.
	// All slices are built on top of arrays.
	// ---
	// A slice is a descriptor for a  contiguous  segment  of
	// an underlying array and provides  access to a numbered
	// sequence of elements from that  array.  A  slice  type
	// denotes the set of all slices of arrays of its element
	// type. The value of an uninitialized slice is nil.
	// https://golang.org/ref/spec#Slice_types

	// This is how we declare a slice
	//myCourses := make([]string, 5, 10)

	// This is a shorter way to do it
	myCourses := []string{"Docker", "Puppet", "Python"}

	fmt.Printf("Length: %d\nCapacity: %d\n", len(myCourses), cap(myCourses))
	fmt.Println(reflect.TypeOf(myCourses))
}
