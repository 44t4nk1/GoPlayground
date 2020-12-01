package main

import (
	"fmt"
	"math"
)

type shapes interface {
	area()
}

type square struct {
	side int
}

type circle struct {
	radius int
}

func (c circle) area() {
	fmt.Println(math.Pi * float64(c.radius*c.radius))
}

func (s square) area() {
	fmt.Println(s.side * s.side)
}

func getArea(s shapes) {
	s.area()
}

func main() {
	c1 := circle{radius: 3}
	s1 := square{side: 3}
	getArea(c1)
	getArea(s1)
}
