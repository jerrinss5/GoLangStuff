package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	// explicity type assignment
	// var b *int = &a
	// implict type assignment to a pointer by compiler
	var b = &a

	fmt.Println(b)
	fmt.Println(*b)

	// assinging the value of 42 to the address location dereferenced by b
	*b = 42
	fmt.Println(a)
}
