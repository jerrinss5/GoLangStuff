package main

import (
	"fmt"
	"time"
)

func main() {
	// making a unbuffered channel to communicate int
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// adding the value of i to the channel
			c <- i
		}
	}()

	go func() {
		for {
			// reading the value of int from the channel
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
