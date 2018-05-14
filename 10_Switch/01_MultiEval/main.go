package main

import "fmt"

func main() {
	switch "test" {
	case "test", "tester":
		fmt.Println("It is either test or tester")
	default:
		fmt.Println("This is nobody")
	}
}
