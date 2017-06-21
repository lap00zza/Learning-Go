package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	_, err := os.Open("./test.txt")
	if err != nil {
		log.Println(err)
	}

	// This will not result in an error.
	// So return value of err is nil.
	_, err1 := os.Open("D:\\Golang Projects\\src\\Learning Go\\src\\func.go")
	fmt.Println("Error: ", err1)

	_, err2 := cool("Is this cool enough?")
	if err2 != nil {
		log.Println(err2)
	}
}

func cool(text string) (string, error) {
	return "", errors.New("Error converting text. Its not cool enough!")
}
