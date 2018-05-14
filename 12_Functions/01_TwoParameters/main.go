package main

import "fmt"

func main() {
	greet("Roronoa", "Zoro")
}

func greet(fname, lname string) {
	fmt.Println(fname, " ", lname)
}
