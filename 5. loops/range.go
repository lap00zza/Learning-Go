package main

import "fmt"

func main()  {

	inProgress := []string{"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}
	completed := []string{"Docker Deep Dive", "Go Fundamentals", "Puppet Fundamentals"}

	for _, i := range inProgress {
		fmt.Println("\n---")
		fmt.Println(i)
		for _, j := range completed {
			if i == j {
				fmt.Println("Clash *-> ", j)
			} else {
				fmt.Println(j)
			}
		}
	}
}
