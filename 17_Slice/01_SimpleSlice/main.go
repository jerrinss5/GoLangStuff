package main

import "fmt"

func main() {
	numberSlice := []string{"1", "2", "3", "4", "5", "6"}
	// printing all the slice values
	fmt.Println(numberSlice)
	// slicing the slice
	fmt.Println(numberSlice[2:4])
	// accessing a value in slice by index
	fmt.Println(numberSlice[2])
	// accessing slice directly
	fmt.Println("string"[2])
}
