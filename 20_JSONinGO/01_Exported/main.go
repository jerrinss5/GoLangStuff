package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person : struct to store persons information
type Person struct {
	First string
	Last  string
	Age   int
	// since it is not exported it won't be marshalled to json
	notExported int
}

func main() {
	p1 := Person{"Roronoa", "Zoro", 28, 007}
	byteSlice, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T \n", byteSlice)
	fmt.Println(string(byteSlice))
}
