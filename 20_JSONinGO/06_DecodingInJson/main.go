package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Person : struct to store persons information
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"Roronoa", "Last":"Zoro", "Age":28}`)
	// decoding a json input
	err := json.NewDecoder(rdr).Decode(&p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Println(p1.notExported)
}
