package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// here the receieve (p person) indicates that this function belongs the struct person
func (p person) fullname() string {
	return p.first + p.last + " : " + string(p.age)
}

func main() {
	p1 := person{"Monkey", "Luffy", 24}
	p2 := person{"Roronoa", "Zoro", 28}
	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
