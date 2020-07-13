package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	p := math.Pi
	cr := float64(c.radius)
	af := cr * cr * p
	a := int(af)
	return a
}

func (r rectangle) rAreaCalculator() int {
	a := r.side1 * r.side2
	return a
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]interface{}, 0, 10)
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		seT := triangle{rand.Intn(10), rand.Intn(10)}
		seC := circle{rand.Intn(10)}
		seR := rectangle{rand.Intn(10), rand.Intn(10)}
		rndn := rand.Intn(2)
		switch {
		case rndn == 0:
			s = append(s, seT)
		case rndn == 1:
			s = append(s, seC)
		default:
			s = append(s, seR)
		}

		fmt.Println(s)
	}

}
