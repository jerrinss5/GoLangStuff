package main

import "fmt"

// Person : Struct to store persons information
type Person struct {
	name string
	age  int
}

func main() {
	// storing the address to the pointer p1
	p1 := &Person{"Luffy", 24}
	fmt.Println(p1)
	// pointer to the type main.Person
	fmt.Printf("%T\n", p1)
	// compiler adds the pointer reference at compile time
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	// fmt.Println(*p1.name)
}
