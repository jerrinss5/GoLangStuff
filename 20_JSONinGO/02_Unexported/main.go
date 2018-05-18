package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person : struct to store persons information
type Person struct {
	// nothing would be exported as all of the fields are in lower case
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"Roronoa", "Zoro", 28}
	fmt.Print("Before: ")
	fmt.Println(p1)
	byteSlice, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("After: ")
	fmt.Println(string(byteSlice))
}
