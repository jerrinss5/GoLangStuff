package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// adding two items to the wait group
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("Foo: ", i)
	}
	// will inform the wait group that function is done
	// wait group count is reduced by 1
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar: ", i)
	}
	wg.Done()
}
