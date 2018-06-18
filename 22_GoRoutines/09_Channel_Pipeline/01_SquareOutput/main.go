package main

import (
	"fmt"
)

func main() {
	// passing the numbers whose square needs to be calculated
	// getting the channel for the same
	c := getSqNum(2, 3)
	// passing the channel to calculate the square
	out := calcSq(c)
	// waiting for the result to be consumed
	fmt.Println(<-out)
	fmt.Println(<-out)
}

// get sq num takes in variadc of integers and returns the channel
func getSqNum(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			// looping through the numbers and adding to the channel
			out <- n
		}
		close(out)
	}()
	return out
}

func calcSq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
