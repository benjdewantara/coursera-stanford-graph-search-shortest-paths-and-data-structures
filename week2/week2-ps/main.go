package main

import (
	"./Graph"
	"fmt"
)

func main() {
	g := Graph.Graph{}
	g.PopulateFromFile("_dcf1d02570e57d23ab526b1e33ba6f12_dijkstraData.txt")
	g.Edges = g.Edges[0:200]

	fmt.Println("Hell on earth")
}
