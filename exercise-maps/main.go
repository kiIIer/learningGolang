package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mmap := make(map[string]int)
	ccounter := strings.Fields(s)
	for _, wwords := range ccounter {
		mmap[wwords] = mmap[wwords] + 1
	}
	return mmap
}

func main() {
	wc.Test(WordCount)
}
