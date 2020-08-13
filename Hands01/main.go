package main

import (
	"fmt"
	"math"
)

// sqare structure to hold side variable
type square struct {
	side float64
}

// calculating area for square
func (s square) area() float64 {
	return s.side * s.side
}

// cirlce structure to hold the radius variable
type cirlce struct {
	radius float64
}

// calculating area of cirlce
func (c cirlce) area() float64 {
	return math.Pi * c.radius * c.radius
}

// shape interface parenting both square and cirlce
type shape interface {
	area() float64
}

// print area is a common method to print are of the shape provided
func printArea(s shape) {
	fmt.Println("Area is: ", s.area())
}

func main() {
	s1 := square{side: 20}
	c1 := cirlce{radius: 10}
	printArea(s1)
	printArea(c1)
}
