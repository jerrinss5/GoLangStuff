package main

import "fmt"

// function wrapper of return type function
func wrapper() func() int {
	x := 0
	// returning function which itself returns an int
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
