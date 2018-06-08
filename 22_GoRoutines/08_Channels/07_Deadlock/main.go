package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// adding to the channel i values in a loop
			c <- i
		}
		close(c)
	}()

	// looping over the range of values of the channel until the channel is closed
	for n := range c {
		fmt.Println(n)
	}
}
