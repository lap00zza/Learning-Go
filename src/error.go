package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
	}

	// This will not result in an error.
	// So return value of err is nil.
	_, err1 := os.Open("D:\\Golang Projects\\src\\Learning Go\\src\\func.go")
	fmt.Println("Error: ", err1)

}
