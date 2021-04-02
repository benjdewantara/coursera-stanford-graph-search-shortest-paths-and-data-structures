package main

import (
	"./DijkstraShortestPath"
	"./Graph"
	"fmt"
)

func main() {
	g := Graph.Graph{}
	g.PopulateFromFile("_dcf1d02570e57d23ab526b1e33ba6f12_dijkstraData.txt")
	g.Edges = g.Edges[0:200]

	d := DijkstraShortestPath.DijkstraShortestPath{
		Graph: g,
	}

	d.InitWithRoot(0)

	fmt.Println("Hell on earth")
}
