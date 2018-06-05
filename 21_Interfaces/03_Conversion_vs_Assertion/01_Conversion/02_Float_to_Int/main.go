package main

import "fmt"

func main() {
	var x = 12
	var y = 12.12345

	// narrowing conversion
	fmt.Println(int(y) + x)
}
