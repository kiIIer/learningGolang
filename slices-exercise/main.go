package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {

		for j := 0; j < dy; j++ {
			o := i
			if o > 127 {
				o = 256 - o
			}
			k := j
			if k > 127 {
				k = 256 - k
			}
			k = k / 2
			o = o / 2
			s[i][j] = uint8(k + o)
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
