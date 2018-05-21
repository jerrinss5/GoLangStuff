package main

import "fmt"

// Square : Square struct
type Square struct {
	side float64
}

func (sq Square) area() float64 {
	return sq.side * sq.side
}

// Shape : interface which requires the function to implement area returning float64
type Shape interface {
	area() float64
}

func info(sh Shape) {
	fmt.Println(sh)
	fmt.Println(sh.area())
}

func main() {
	s := Square{10}
	info(s)
}
