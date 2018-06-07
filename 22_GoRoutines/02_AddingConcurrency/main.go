package main

import "fmt"

// this function would not give any results as
// the main, foo and bar are independent routines spooled off
// as soon foo and bar routine starts up the main routine finishes
func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("Foo: ", i)
	}
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar: ", i)
	}
}
