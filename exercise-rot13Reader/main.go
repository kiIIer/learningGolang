package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < n; i++ {
		switch {
		case p[i] < 78 && p[i] > 64:
			p[i] = p[i] + 13
		case p[i] > 77 && p[i] < 91:
			p[i] = p[i] - 13

		case p[i] < 110 && p[i] > 96:
			p[i] = p[i] + 13
		case p[i] < 123 && p[i] > 109:
			p[i] = p[i] - 13
		default:
			p[i] = 32
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
