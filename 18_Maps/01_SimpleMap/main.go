package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["k1"] = 1
	myMap["k2"] = 2

	fmt.Println(myMap)

	v, ok := myMap["k1"]
	fmt.Println("ok?: ", ok, v)

	// can also be declared as
	newMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(newMap)
}
