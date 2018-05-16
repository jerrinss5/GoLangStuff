package main

import "fmt"

func main() {
	// passing some values directly
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3}
	// reading values off the slice and passing it as individual values to the foo function
	foo(aSlice...)
	foo()
}

// foo function which takes in parameter of int declared as variadic
func foo(numbers ...int) {
	for _, n := range numbers {
		fmt.Println(n)
	}
}
