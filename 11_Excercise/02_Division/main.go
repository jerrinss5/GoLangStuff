package main

import "fmt"

func main() {
	var largeNumber int
	var smallNumber int

	fmt.Print("Enter a large number: ")
	fmt.Scan(&largeNumber)

	fmt.Print("Enter a small number: ")
	fmt.Scan(&smallNumber)

	fmt.Println("Result after division: ", largeNumber/smallNumber)
}
