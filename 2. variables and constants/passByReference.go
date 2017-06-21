package main

import "fmt"

func main() {
	name := "Jewel Mahanta"
	course := "Learning Go"

	fmt.Println("\nName: ", name, "\nCourse: ", course)
	changeCourseRef(&course)
	fmt.Println("Changed course: ", course)
}

// courseStr *string means that courseStr is a pointer to a
// string value.
func changeCourseRef(courseStr *string) string {
	*courseStr = "Advanced Go! Concepts"
	fmt.Println("Trying to change course to: ", *courseStr)
	return *courseStr
}
