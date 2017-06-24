package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitgrp sync.WaitGroup
	chan1 := make(chan int)
	waitgrp.Add(2)

	go func() {
		defer waitgrp.Done()
		chan1 <- 2
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("Tx ended")
	}()

	go func() {
		defer waitgrp.Done()
		recv := <-chan1
		fmt.Println(recv)
		fmt.Println("Rx ended")
	}()

	waitgrp.Wait()
}
