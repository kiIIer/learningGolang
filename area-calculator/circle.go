package main

import "math"

//Circle is circle
type Circle struct {
	radius float64
}

func (c Circle) calculateArea() float64 {
	return c.radius * c.radius * math.Pi
}
