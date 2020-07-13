package main

// Rectangle is a rectangle.
type Rectangle struct {
	side1, side2 float64
}

func (r Rectangle) calculateArea() float64 {
	return r.side1 * r.side2
}
