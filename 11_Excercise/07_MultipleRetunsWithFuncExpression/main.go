package main

import "fmt"

func main() {
	// declaring a function expresion which stores the anonymous function
	var doSomething = func(n int) (int, bool) {
		return n / 2, (n%2 == 0)
	}
	// declaring an input variable in an idiomatic way
	var input int
	fmt.Print("Enter a vaule: ")
	// getting the user input
	fmt.Scan(&input)
	// calling the function expression with the user inputted value
	var divisionValue, even = doSomething(input)
	// printing the return results
	fmt.Println(divisionValue)
	fmt.Println(even)
}
