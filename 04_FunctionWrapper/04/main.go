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
	// calling the wrapper function and assiging the function to a variable increment
	increment := wrapper()
	// calling the increment function twice
	fmt.Println(increment())
	fmt.Println(increment())
}
