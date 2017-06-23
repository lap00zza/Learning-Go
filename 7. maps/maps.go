package main

import "fmt"

func main() {

	// Maps are similar to arrays/slices but:
	// 1. they are unordered
	// 2. they are <key>:<value> pairs
	// 3. they are dynamically resizeable
	// 4. they are references
	//
	// +-- https://golang.org/ref/spec#Map_types --+
	// A map is an unordered group of elements of one type, called the
	// element type, indexed by a set of unique keys of another type,
	// called the key type. The value of an uninitialized map is nil.
	//
	// MapType     = "map" "[" KeyType "]" ElementType .
	// KeyType     = Type .
	//
	// The comparison operators == and != must be fully defined for
	// operands of the key type; thus the key type must not be a
	// function, map, or slice. If the key type is an interface type,
	// these comparison operators must be defined for the dynamic key
	// values; failure will cause a run-time panic.

	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	// This syntax, where we declare and initialize struct, array,
	// slice, and maps is also know as composite literals.
	// https://golang.org/ref/spec#Composite_literals
	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}
	fmt.Println(leagueTitles, recentHead2Head)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	for key, val := range testMap {
		// %v the value in a default format when printing
		// structs, the plus flag (%+v) adds field names
		fmt.Printf("\nKey: %v | Value: %v", key, val)
	}

	mySlice := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}
	fmt.Printf("\n\n%v | A: %v", mySlice, mySlice["A"])

	// How this works is, if this key is already present
	// its associated value gets updated. If the key does
	// not exist, a new key gets created with the value.
	mySlice["A"] = 100
	mySlice["New Key"] = 200
	fmt.Printf("\n%v", mySlice)

	// To delete a key, we use the built-in function 'delete'
	delete(mySlice, "New Key")
	fmt.Printf("\n%v", mySlice)
}
