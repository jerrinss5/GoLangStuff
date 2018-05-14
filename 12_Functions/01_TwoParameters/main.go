package main

import "fmt"

func main() {
	// arguments passed to the function
	greet("Roronoa", "Zoro")
}

// functions declared with parameters
func greet(fname, lname string) {
	fmt.Println(fname, " ", lname)
}
