package main

import "fmt"

// declaring a constant at package scope level
const outside string = "This is an outside varialbe"

func main() {
	// declaring a constant at the function scope level
	const inside = "This is an inside variable"
	fmt.Println("Inside variable - ", inside)
	fmt.Println("Outside variable - ", outside)
}
