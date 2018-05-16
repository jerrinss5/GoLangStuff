package main

import "fmt"

func main() {

	greeeting := []string{
		"test",
		"tester",
		"testing",
		"tested",
		"test_test",
		"test_load",
		"test_loader",
	}

	// get elements from 1 uptil 2 - 1
	fmt.Print("[1:2] - ")
	fmt.Println(greeeting[1:2])

	// get elements uptil 2 - 0 1
	fmt.Print("[:2] - ")
	fmt.Println(greeeting[:2])

	// get elements from 5 uptil end - 5 6
	fmt.Print("[5: ] - ")
	fmt.Println(greeeting[5:])

	// get everything
	fmt.Print("[:] - ")
	fmt.Println(greeeting[:])
}
