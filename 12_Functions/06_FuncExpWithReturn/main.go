package main

import "fmt"

// declaring a function with return type function of type string
func greeting() func() string {
	// return value of type function of type string
	return func() string {
		return "Greetings Hooman!!!!"
	}
}

func main() {
	// gettung the function expression
	greet := greeting()

	// calling the function expression
	fmt.Println(greet())
}
