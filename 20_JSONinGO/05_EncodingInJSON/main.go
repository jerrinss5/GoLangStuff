package main

import (
	"encoding/json"
	"log"
	"os"
)

// Person : Struct to store persons information
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Roronoa", "Zoro", 28, 007}
	// encoding the person structure JSON to stdout
	err := json.NewEncoder(os.Stdout).Encode(p1)
	if err != nil {
		log.Fatal(err)
	}
}
