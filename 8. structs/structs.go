package main

import (
	"fmt"
	"reflect"
)

func main() {

	// A struct is a sequence of named elements, called fields,
	// each of which has a name and a type. Field names may be
	// specified explicitly (IdentifierList) or implicitly
	// (AnonymousField). Within a struct, non-blank field names
	// must be unique.
	// https://golang.org/ref/spec#Struct_types

	// We define a custom type
	// https://golang.org/ref/spec#Types
	type CourseMeta struct {
		Name   string
		Author string
		Level  string
		Rating float64
	}

	// Method 1
	//var course CourseMeta

	// Method 2
	// using 'new' gives us a pointer
	//course := new(CourseMeta)

	// Method 3
	// For struct literals the following rules apply:
	// 1. A key must be a field name declared in the struct type.
	// 2. An element list that does not contain any keys must list
	//	  an element for each struct field in the order in which
	//    the fields are declared.
	// 3. If any element has a key, every element must have a key.
	// 4. An element list that contains keys does not need to have
	//    an element for each struct field. Omitted fields get the
	//    zero value for that field.
	// 5. A literal may omit the element list; such a literal
	//   evaluates to the zero value for its type.
	//
	// TODO: understand what the point below means
	// * It is an error to specify an element for a non-exported
	//   field of a struct belonging to a different package. [x]

	// Point 1
	course1 := CourseMeta{
		Author: "Nigel Paulton",
		Level:  "Beginner",
		Name:   "Docker Deep Dive",
		Rating: 5,
	}

	// Point 2
	course2 := CourseMeta{"Go! in the Web", "Jewel Mahanta", "Beginner", 4.1}

	// Point 3
	// Results in a error
	//course3 := CourseMeta{
	//	Author: "Nigel Paulton",
	//	"Beginner",
	//	"Docker Deep Dive",
	//	5,
	//}

	// point 4
	course4 := CourseMeta{
		Author: "Mickey Mouse",
	}

	// point 5
	course5 := CourseMeta{}

	// Results
	fmt.Printf(
		"\n%v\n%v\ncourse4.Rating: %v | type: %v\ncourse5.Rating: %v | type: %v",
		course1, course2, course4.Rating, reflect.TypeOf(course4.Rating), course5.Rating,
		reflect.TypeOf(course5.Rating),
	)
}
