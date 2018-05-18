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
}

func main() {
	var p1 Person
	fmt.Println("Uninitialized:")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"Roronoa", "Last":"Zoro", "Age":28}`)
	// taking the byte slice of JSON input and adding it to the address referenced by p1
	err := json.Unmarshal(bs, &p1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After Unmarshalling:")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T\n", p1)
}
