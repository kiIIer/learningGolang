package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walk2(t, ch)
	close(ch)
}

//Walk2 is like walk
func Walk2(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk2(t.Left, ch)
		ch <- t.Value
		Walk2(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		ct1, ok1 := <-ch1
		ct2, ok2 := <-ch2
		if ok1 != ok2 || ct1 != ct2 {
			return false
		}
		if !ok1 {
			break
		}

	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
