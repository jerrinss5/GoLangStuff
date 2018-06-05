package main

import "fmt"

func main() {
	var x = 12
	var y = 12.12345
	// widening the type
	fmt.Println(y + float64(x))
}
