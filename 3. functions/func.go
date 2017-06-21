package main

import (
	"fmt"
	"strings"
)

func main() {

	module := "function basics"
	name := "jewel mahanta"
	fmt.Println(converter(module, name))
}

func converter(module, name string) (string, string) {
	module = strings.ToUpper(module)
	name = strings.Title(name)
	return module, name
}
