package main

import "fmt"

func main() {
	// idiomatic declaration of a variable inside function
	var input int
	fmt.Print("Enter a number greater than zero: ")
	// getting the input from the user
	fmt.Scan(&input)
	// calling the function which returns 2 values
	var division, even = doSomething(input)
	fmt.Println(division)
	fmt.Println(even)
}

// function with two parameters
func doSomething(n int) (int, bool) {
	// check the value after its division by 2 and if its an even number
	return n / 2, (n%2 == 0)
}
