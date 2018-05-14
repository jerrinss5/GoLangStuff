package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)
	// passing the address value to the change fucntion
	change(&a)
	fmt.Println(a)
}

// accepting the address as a pointer data type
func change(a *int) {
	// changing the address pointer data type to a new value
	*a = 50
	fmt.Println(&a)
}
