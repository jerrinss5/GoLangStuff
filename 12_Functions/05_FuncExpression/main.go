package main

import "fmt"

func main() {
	// anonymous function assigned to greeting as func expression
	greeting := func() {
		fmt.Println("Greetings Hooman!!!!")
	}

	// using the func expression to call the anonymous function
	greeting()

	// checking the type of the function
	fmt.Printf("%T\n", greeting)
}
