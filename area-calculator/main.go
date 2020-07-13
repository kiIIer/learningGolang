package main

import (
	"fmt"
)

type triangle struct {
	side1, side2 int
}

type circle struct {
	radius int
}

type rectangle struct {
	side1, side2 int
}

func (t triangle) tAreaCalculator() int {
	a := t.side1 * t.side2 / 2
	return a
}

func (c circle) cAreaCalculator() int {
	p := 3
	a := c.radius * c.radius * p
	return a
}

func (r rectangle) rAreaCalculator() int {
	a := r.side1 * r.side2
	return a
}

func main() {
	t := triangle{4, 5}
	c := circle{10}
	r := rectangle{6, 8}

	fmt.Println(t)
	fmt.Println(t.tAreaCalculator())
	fmt.Println(c)
	fmt.Println(c.cAreaCalculator())
	fmt.Println(r)
	fmt.Println(r.rAreaCalculator())
}
