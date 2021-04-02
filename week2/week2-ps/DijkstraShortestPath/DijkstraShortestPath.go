package DijkstraShortestPath

import "../Graph"
import "../Heap"

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

}
