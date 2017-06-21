package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice[4], len(mySlice), cap(mySlice))

	mySlice[4] = 99
	fmt.Println(mySlice)

	// When we use [0:2] syntax, in [1,2,3,4,5] we will get
	// [1,2,3]. It means slice from 0 to 2. 0 is included
	// and 2 is excluded in the final slice.
	// Remember:
	// [:4] -> means 0 to 4
	// [4:] -> means 4 to length of array - 1
	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)

	// append generates a new slice if cap is not sufficient
	mySlice = append(mySlice, 12, 13, 14)
	fmt.Println(mySlice, len(mySlice), cap(mySlice))
}
