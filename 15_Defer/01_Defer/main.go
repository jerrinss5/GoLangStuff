package main

import "fmt"

func world() {
	fmt.Println("World")
}

func hello() {
	fmt.Println("Hello")
}

func main() {
	// defers the world function till the end of the main function before exiting it
	defer world()
	hello()
}
