package main

import "fmt"

func main() {
	name := "Jewel Mahanta"
	course := "Learning Go"

	fmt.Println("\nName: ", name, "\nCourse: ", course)
	changeCourse(course)
	fmt.Println("Changed course: ", course)
}

func changeCourse(course string) string {
	course = "Advanced Go! Concepts"
	fmt.Println("Trying to change course to: ", course)
	return course
}
