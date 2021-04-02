package DijkstraShortestPath

import "../Graph"

type DijkstraShortestPath struct {
	Graph               Graph.Graph
	VerticesCongregated []int
	WeightsCongregated  []int
}

func (d *DijkstraShortestPath) InitWithRoot(initialRootIndex int) {
	edgeOfRoot := d.Graph.Edges[initialRootIndex]
	d.VerticesCongregated = append(d.VerticesCongregated, edgeOfRoot.Tail)
	d.WeightsCongregated = append(d.WeightsCongregated, 0)
}
