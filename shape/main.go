package main

import (
	"fmt"
	"math"
)

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	sa := square{sideLength: 10}
	ta := triangle{
		height: 10,
		base:   10,
	}
	printArea(sa)
	printArea(ta)
}

func printArea(s shape) {
	fmt.Printf("Area is : %v\n", s.getArea())
}
func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
