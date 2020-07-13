package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]interface{}, 0, 10)
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		seT := Triangle{float64(rand.Intn(10) + 1), float64(rand.Intn(10) + 1)}
		seC := Circle{float64(rand.Intn(10) + 1)}
		seR := Rectangle{float64(rand.Intn(10) + 1), float64(rand.Intn(10) + 1)}
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
	for i := 0; i < len(s); i++ {
		ct := s[i]
		t, okT := ct.(Triangle)
		c, okC := ct.(Circle)
		r, _ := ct.(Rectangle)
		switch {
		case okT:
			a := t.calculateArea()
			fmt.Printf("%+v area: %.2f\n", ct, a)
		case okC:
			a := c.calculateArea()
			fmt.Printf("%+v area: %.2f\n", ct, a)
		default:
			a := r.calculateArea()
			fmt.Printf("%+v area: %.2f\n", ct, a)
		}

	}
}
