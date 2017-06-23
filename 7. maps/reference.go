package main

import "fmt"

func main()  {

	myArr := [4]int{1, 2, 3, 4}
	mySlice := []int{1, 2, 3, 4}
	myMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}

	fmt.Println(myArr, mySlice, myMap)

	// Arrays and Slices are passed by value
	awesomeArr(myArr)
	fmt.Println("myArr: ", myArr)
	awesomeSlice(mySlice)
	fmt.Println("mySlice: ", mySlice)

	// NOTE:
	// Maps are passed by reference
	// Changes made to map are visible to caller.
	// Maps are not safe for concurrency.
	// Maps are cheap to pass.
	awesomeMap(myMap)
	fmt.Println("myMap: ", myMap)
}

func awesomeArr(arg [4]int)  {
	arg[3] = 500
}

func awesomeSlice(arg []int)  {
	arg = append(arg, 5, 6, 7)
}

func awesomeMap(arg map[string]int)  {
	arg["Cool"] = 100
}