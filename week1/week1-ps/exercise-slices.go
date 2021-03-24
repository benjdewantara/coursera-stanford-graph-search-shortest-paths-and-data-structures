package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var p = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
	}

	for yIdx := 0; yIdx < dy; yIdx++ {
		for xIdx := 0; xIdx < dx; xIdx++ {
			//p[yIdx][xIdx] = uint8((xIdx + yIdx) / 2)
			p[yIdx][xIdx] = uint8(xIdx * yIdx)
			//p[yIdx][xIdx] = uint8(xIdx ^ yIdx)
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}
