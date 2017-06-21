package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Runtime ", runtime.GOOS)
}
