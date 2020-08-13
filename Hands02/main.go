package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) pSpeak() {
	fmt.Println(p.fname, p.lname, "here!")
}

type secretAgent struct {
	p             person
	licenseToKill bool
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.p.fname, sa.p.lname, "at your service!")
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

	p.pSpeak()
	sa.saSpeak()
	sa.p.pSpeak()
}
