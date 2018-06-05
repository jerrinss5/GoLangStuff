package main

import "fmt"

func main() {
	var name interface{} = "Sydney"

	str, ok := name.(string)

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Value is not a string")
	}
}
