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

func (d *DijkstraShortestPath) CongregateOnce() {
	node, weight := d.MinHeapOfVerticesAdjacent.ExtractRoot()

	d.VerticesCongregated = append(d.VerticesCongregated, node)
	d.WeightsCongregated = append(d.WeightsCongregated, weight)

	edge := d.Graph.GetEdgeAt(node)
	heads := edge.Heads

	for headNodeIndx, headNode := range heads {
		if indexInIntegerArray(d.VerticesCongregated, headNode) >= 0 {
			continue
		}

		edge.Weights[headNodeIndx] += weight

		nextNode := edge.Heads[headNodeIndx]
		nextWeight := edge.Weights[headNodeIndx]
		d.MinHeapOfVerticesAdjacent.Insert(nextNode, nextWeight)
	}
}

func indexInIntegerArray(arr []int, elem int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return i
		}
	}

	return -1
}
