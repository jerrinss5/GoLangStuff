package main

import (
	"fmt"
)

func main() {
	// passing the value whose factorial needs to be calculated
	c1 := factorial(4)
	// getting the channel for the factorial function to factorial calculator
	c2 := factorialCalc(c1)
	fmt.Println("Value of factorial is:", <-c2)
}

// create a factorial puller which returns the channel of int type
func factorial(n int) chan int {
	// create a channel of type int
	out := make(chan int)
	// making a go routine which loops through and returns numbers in decreasing fashion
	// this runs off independently and program jumps to the return statement
	go func() {
		for i := n; i > 0; i-- {
			// putting the number on to the channel
			out <- i
		}
		// closing the channel
		close(out)
	}()
	// returning the channel to the caller
	return out
}

// a factorial receiver which receives the channel of the puller and returns the output
func factorialCalc(c chan int) chan int {
	out := make(chan int)
	// go routine to calculate the factorial
	go func() {
		total := 1
		for n := range c {
			total *= n
		}
		out <- total
		close(out)
	}()
	return out
}
