package main

import (
	"fmt"
)

func main() {
	// making channel to receive int
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// when channel is closed you cannot put values to the channel
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	// no need to add sleep as the range c would wait to receive the values off of the channel
}
