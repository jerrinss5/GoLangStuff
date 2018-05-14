package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	// assigning the address location of a to pointer int type in b
	var b *int = &a

	fmt.Println(b)
}
