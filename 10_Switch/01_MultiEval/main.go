package main

import "fmt"

func main() {
	switch "test" {
	// checking for test or tester
	case "test", "tester":
		fmt.Println("It is either test or tester")
	default:
		fmt.Println("This is nobody")
	}
}
