package main

import "fmt"

// Person : structure which stores a persons information
type Person struct {
	First string
	Last  string
	Age   int
}

// DoubleZero : structure which stores an agent information with kill contract
type DoubleZero struct {
	// here person is the anonymous field
	Person
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Monkey",
			Last:  "Luffy",
			Age:   24,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Roronoa",
			Last:  "Zoro",
			Age:   29,
		},
		LicenseToKill: true,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
}
