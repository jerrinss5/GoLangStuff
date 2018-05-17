package main

import "fmt"

func main() {
	// with make
	var myMap = make(map[string]string)
	fmt.Println(myMap == nil)

	myMap["hello"] = "world"
	myMap["bye"] = "bye"

	fmt.Println(myMap)

	// just declaration
	var myMap1 map[string]string
	fmt.Println(myMap1 == nil)

	// with composite literal
	var myMap2 = map[string]string{}
	fmt.Println(myMap2 == nil)
}
