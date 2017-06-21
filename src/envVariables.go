package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Environ())
	envVars := os.Environ()
	for i := 0; i < len(envVars); i++ {
		fmt.Println(envVars[i])
	}
}
