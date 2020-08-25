package main

import (
	"fmt"
)

func caesarCipher(s string, k int) string {
	const min, max = 'a', 'z'
	const count = max - min + 1
	const miN, maX = 'A', 'Z'

	rotator := rune(k % count)
	ret := ""
	codeS := make([]rune, len(s))
	for i, ch := range s {
		if ch >= min && ch <= max {
			ch += rotator

			if ch > max {
				ch -= count

			}
		}
		if ch >= miN && ch <= maX {
			ch += rotator
			if ch > maX {
				ch -= count
			}
		}
		codeS[i] = ch
		ret = string(codeS)
	}
	return ret
}

func main() {
	var lenght, rot int
	var input string
	fmt.Scanf("%d\n", &lenght)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &rot)
	done := caesarCipher(input, rot)
	fmt.Println(done)
}
