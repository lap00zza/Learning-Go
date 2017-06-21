package main

import "fmt"

func main() {
iLoop:
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
		for j := 0; j < 10; j++ {
			fmt.Println("j: ", j)
			for k := 0; k < 10; k++ {
				if i == 5 && j == 5 && k == 5 {
					fmt.Println("Breaking, i: ", i, " j: ", j, " k: ", k)
					break iLoop
				}
				fmt.Println("k: ", k)
			}
		}
	}
}
