package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1
	for _, ch := range input {
		min, max := 'A', 'Z'
		if ch >= min && ch <= max {
			answer++
		}
	}
	fmt.Println(answer)
}
