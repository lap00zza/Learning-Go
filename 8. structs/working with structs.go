package main

import "fmt"

func main() {

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	p := Person{
		firstName: "Jewel",
		lastName:  "Mahanta",
		age:       21,
	}

	// We can access individual struct fields using the
	// period (.) operator. Below we are accessing the
	// field firstName.
	fmt.Println(p.firstName)

	// We can also change the value of a field using the
	// period (.) operator.
	p.age = 200
	fmt.Println(p)
}
