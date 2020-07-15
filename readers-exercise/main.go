package main

import "golang.org/x/tour/reader"

//MyReader is mine
type MyReader struct{}

func (r MyReader) Read(a []byte) (int, error) {
	for i := range a {
		a[i] = 'A'
	}
	return len(a), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
