package main

import "fmt"

func main() {
	fmt.Println(greet("Monkey ", "Luffy"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
