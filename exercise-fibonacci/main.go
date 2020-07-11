package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	number1 := 0
	number2 := 1
	return func() int {
		number := number1
		number1 = number2
		number2 = number + number1
		return number
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
