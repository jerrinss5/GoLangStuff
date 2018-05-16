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

// defere us usually used to close files or other resources, but in code placed just after opening them and then defering to till the end
