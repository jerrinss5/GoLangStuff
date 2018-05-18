package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person : struct to hold persons information
type Person struct {
	First string
	// json:"-" indicates not be marshalled
	Last string `json:"-"`
	// json:"<value> indicates it needs to exported as the value specified"
	Age int `json:"wisdom-score"`
}

func main() {
	// shorthand assignemnt of the struct
	p1 := Person{"Luffy", "Monkey", 24}
	fmt.Print("Before: ")
	fmt.Println(p1)

	// marshalling the struct
	byteSlice, err := json.Marshal(p1)
	// if err handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("After: ")
	fmt.Println(string(byteSlice))
}
