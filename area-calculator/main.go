package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type triangle struct {
	width, hight float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	side1, side2 float64
}

func (t triangle) calculateArea() float64 {
	return t.width * t.hight / 2
}
func (c circle) calculateArea() float64 {
	return c.radius * c.radius * math.Pi
}
func (r rectangle) calculateArea() float64 {
	return r.side1 * r.side2
}

type areaCalculator interface {
	calculateArea() float64
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]areaCalculator, 0, 10)
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		seT := triangle{float64(rand.Intn(10) + 1), float64(rand.Intn(10) + 1)}
		seC := circle{float64(rand.Intn(10) + 1)}
		seR := rectangle{float64(rand.Intn(10) + 1), float64(rand.Intn(10) + 1)}
		rndn := rand.Intn(3)
		switch {
		case rndn == 0:
			s = append(s, seT)
		case rndn == 1:
			s = append(s, seC)
		default:
			s = append(s, seR)
		}

	}
	fmt.Printf("%+v\n", s)
	for i := range s {
		ct := s[i]
		a := ct.calculateArea()
		fmt.Printf("%+v area: %.2f\n", ct, a)

	}
}
