package main

import (
	"./DijkstraShortestPath"
	"./Graph"
	"./Utility"
	"fmt"
	"strings"
)

func main() {
	g := Graph.Graph{}
	g.PopulateFromFile("_dcf1d02570e57d23ab526b1e33ba6f12_dijkstraData.txt")
	g.Edges = g.Edges[0:200]

	d := DijkstraShortestPath.DijkstraShortestPath{
		Graph: g,
	}

	d.InitWithRoot(0)

	for d.HasAdjacentVertices() {
		d.CongregateOnce()
	}

	resultOneline := ""

	verticesOfInterest := []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}
	for _, vert := range verticesOfInterest {
		indx := Utility.IntArr(d.VerticesCongregated).IndexOf(vert)

		resultOneline = fmt.Sprintf("%s,%d", resultOneline, d.WeightsCongregated[indx])
		fmt.Println(fmt.Sprintf("node=%d distance=%d", d.VerticesCongregated[indx], d.WeightsCongregated[indx]))
	}

	resultOneline = strings.TrimPrefix(resultOneline, ",")

	fmt.Println(resultOneline)
}
