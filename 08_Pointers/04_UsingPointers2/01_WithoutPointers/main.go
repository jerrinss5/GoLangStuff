package main

import "fmt"

func main() {
	x := 5
	// passing the value 5
	change(x)
	// value of x here is unchanged
	fmt.Println(x)
}

func change(x int) {
	// value 5 is ignored and x is assigned a value which is block scope
	x = 0
}
