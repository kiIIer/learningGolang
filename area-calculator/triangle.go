package main

//Triangle is triangle
type Triangle struct {
	width, hight float64
}

func (t Triangle) calculateArea() float64 {
	return t.width * t.hight / 2
}