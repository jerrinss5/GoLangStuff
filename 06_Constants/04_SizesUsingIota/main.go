package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10) // 01 shifting 10 times to the left would be 1KB
	mb = 1 << (iota * 10) // 1 << (2 * 10)
	gb = 1 << (iota * 10) // 2 << (3 * 10)
)

func main() {
	fmt.Println("Binary\tDecimal")
	// printing binary and decimal representation of KB
	fmt.Printf("%b\t", kb)
	fmt.Printf("%d\n", kb)

	// printing binary and decimal representation of MB
	fmt.Printf("%b\t", mb)
	fmt.Printf("%d\n", mb)

	// printing binary and decimal representation of GB
	fmt.Printf("%b\t", gb)
	fmt.Printf("%d\n", gb)
}
