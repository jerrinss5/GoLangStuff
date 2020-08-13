package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) speak() string {
	return fmt.Sprintln(p.fname, p.lname, "here!")
}

type secretAgent struct {
	p             person
	licenseToKill bool
}

func (sa secretAgent) speak() string {
	return fmt.Sprintln(sa.p.fname, sa.p.lname, "at your service!")
}

type agent interface {
	speak() string
}

func speaking(a agent) {
	fmt.Println(a.speak())
}

func main() {
	p := person{
		fname: "Monkey",
		lname: "Luffy",
	}
	sa := secretAgent{
		p: person{
			fname: "Roronoa",
			lname: "Zoro",
		},
		licenseToKill: true,
	}
	speaking(p)
	speaking(sa)
}
