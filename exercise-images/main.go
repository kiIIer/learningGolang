package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

//Image is image
type Image struct {
}

//ColorModel is ColorModel
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

//Bounds are bounds
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

//At is At
func (img Image) At(x int, y int) color.Color {
	v := uint8(x + y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
