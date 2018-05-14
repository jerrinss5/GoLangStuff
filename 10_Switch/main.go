package main

import "fmt"

func main() {
	switch "test1" {
	case "test":
		fmt.Println("test was here")
		// if fall through is mentioned it would fall through to the next case
		// fallthrough
	case "tester":
		fmt.Println("tester was here")
	default:
		fmt.Println("No one was here")
	}
}
