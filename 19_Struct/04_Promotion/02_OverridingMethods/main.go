package main

import "fmt"

// Person : structure to store persons details
type Person struct {
	First string
	Last  string
	Age   int
}

// DoubleZero : structure to store agents information
type DoubleZero struct {
	Person
	LicenseToKill bool
}

// Greeting : Pirate king
func (p Person) Greeting() string {
	return "be the Pirate King"
}

// Greeting : Greatest Swordsmen
// Overriding the Greeting method associated to person structure
func (doublzero DoubleZero) Greeting() string {
	return "be the greatest Swordsmen in the All seas"
}

func main() {
	p1 := Person{
		First: "Monkey D.",
		Last:  "Luffy",
		Age:   24,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Roronoa",
			Last:  "Zoro",
			Age:   29,
		},
		LicenseToKill: true,
	}

	fmt.Println(p1.First, p1.Last, "will", p1.Greeting())
	fmt.Println(p2.First, p2.Last, "will not", p2.Person.Greeting())
	fmt.Println(p2.First, p2.Last, "will", p2.Greeting())
}
