package main

import (
	"fmt"
	"math"
)

// Square : Square struct
type Square struct {
	side float64
}

// Circle : Circle struct
type Circle struct {
	radius float64
}

func (sq Square) area() float64 {
	return sq.side * sq.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
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
	// passing the square variable - since the square implements area
	// info accepts the square variable
	info(s)

	c := Circle{3}
	info(c)
}
