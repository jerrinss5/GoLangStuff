package main

import "fmt"

// Person : Structure to store persons information
type Person struct {
	First string
	Last  string
	Age   int
}

// DoubleZero : Structure to store agent information with kill licencse
type DoubleZero struct {
	Person
	// overriding the First field in the agent struct
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Monkey",
			Last:  "Luffy",
			Age:   24,
		},
		First:         "Pirate King",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Roronoa",
			Last:  "Zoro",
			Age:   29,
		},
		First:         "Pirate Hunter",
		LicenseToKill: true,
	}

	fmt.Println(p1.Person.First, p1.Last, p1.Age, p1.First)
	fmt.Println(p2.Person.First, p2.Last, p2.Age, p2.First)
}
