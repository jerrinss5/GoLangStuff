package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota // 0
	e = iota // 1
	f = iota // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
