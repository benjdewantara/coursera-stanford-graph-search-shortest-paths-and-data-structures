package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	panic("implement me")
}

func (img Image) Bounds() image.Rectangle {
	panic("implement me")
}

func (img Image) At(x, y int) color.Color {

	panic("implement me")
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
