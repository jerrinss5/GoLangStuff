package main

import "fmt"

func main() {
	a := 42
	// printing the value of a
	fmt.Println(a)
	// printing the memory address of a
	fmt.Println(&a)

	// assigning memory address of a to b int pointer type
	var b *int = &a

	// printing the value of b - memory address of a
	fmt.Println(b)

	// printing the value of memory address of b
	fmt.Println(*b)
}
