package DijkstraShortestPath

import (
	"../Graph"
	"../Heap"
	"../Utility"
)

type DijkstraShortestPath struct {
	Graph                     Graph.Graph
	VerticesCongregated       []int
	WeightsCongregated        []int
	MinHeapOfVerticesAdjacent Heap.MinHeap
}

func (d *DijkstraShortestPath) InitWithRoot(initialRootIndex int) {
	edgeOfRoot := d.Graph.Edges[initialRootIndex]
	d.VerticesCongregated = append(d.VerticesCongregated, edgeOfRoot.Tail)
	d.WeightsCongregated = append(d.WeightsCongregated, 0)

	d.MinHeapOfVerticesAdjacent = Heap.MinHeap{
		Indices: edgeOfRoot.Heads[0:],
		Values:  edgeOfRoot.Weights[0:],
	}
	d.MinHeapOfVerticesAdjacent.Heapify()
}

func (d *DijkstraShortestPath) CongregateOnce() {
	node, weight := d.MinHeapOfVerticesAdjacent.ExtractRoot()

	d.VerticesCongregated = append(d.VerticesCongregated, node)
	d.WeightsCongregated = append(d.WeightsCongregated, weight)

	edge := d.Graph.GetEdgeAt(node)

	for headNodeIndx, headNode := range edge.Heads {
		if Utility.IntArr(d.VerticesCongregated).IndexOf(headNode) >= 0 {
			continue
		}

		edge.Weights[headNodeIndx] += weight

		nextNode := edge.Heads[headNodeIndx]
		nextWeight := edge.Weights[headNodeIndx]
		d.MinHeapOfVerticesAdjacent.UpdateIfLesser(nextNode, nextWeight)
	}
}

func (d *DijkstraShortestPath) HasAdjacentVertices() bool {
	return len(d.MinHeapOfVerticesAdjacent.Values) > 0
}
