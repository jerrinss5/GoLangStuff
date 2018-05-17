package main

import "fmt"

func main() {
	// map declaration and assignment
	myMap := map[int]string{
		0:  "luffy",
		1:  "zoro",
		2:  "sanji",
		3:  "nami",
		4:  "usopp",
		5:  "frank",
		6:  "chopper",
		7:  "brook",
		8:  "robin",
		9:  "chopper",
		10: "jimbei",
	}
	// looping over the range of map values
	for key, val := range myMap {
		fmt.Println(key, " - ", val)
	}
}
