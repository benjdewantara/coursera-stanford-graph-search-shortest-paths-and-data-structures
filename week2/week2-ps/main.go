package main

import (
	"./Heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type WeightedEdge struct {
	tail   int
	heads  []int
	weight []int
}

type Graph struct {
	edges []WeightedEdge
}

type DijkstraShortestPath struct {
	X []int
}

func (g *Graph) populateFromFile(filepath string) {
	contentBytes, _ := ioutil.ReadFile(filepath)
	for _, edgeStr := range strings.Split(string(contentBytes), "\n") {
		edgeStr = strings.TrimSuffix(edgeStr, " ")
		edgeStrSplit := strings.Split(edgeStr, "\t")
		tail, _ := strconv.Atoi(edgeStrSplit[0])
		e := WeightedEdge{
			tail: tail,
		}
		for _, headWeightStr := range edgeStrSplit[1:] {
			headWeightStrSplit := strings.Split(headWeightStr, ",")
			if len(headWeightStrSplit) != 2 {
				continue
			}
			head, _ := strconv.Atoi(headWeightStrSplit[0])
			weight, _ := strconv.Atoi(headWeightStrSplit[1])
			e.heads = append(e.heads, head)
			e.weight = append(e.weight, weight)
		}

		g.edges = append(g.edges, e)
	}
}

func main() {
	//g := Graph{}
	//g.populateFromFile("_dcf1d02570e57d23ab526b1e33ba6f12_dijkstraData.txt")
	//g.edges = g.edges[0:200]

	MAX_LENGTH := 5
	num := 5

	h := Heap.MinHeap{}

	for i := 0; i < MAX_LENGTH; i++ {
		h.Indices = append(h.Indices, i)
		//h.Indices[i] = i

		h.Values = append(h.Values, num)
		//h.Values[i] = num

		num--
	}

	//order0 := 1<<0
	//fmt.Println(order0)
	//
	//order1 := 1<<1
	//fmt.Println(order1)
	//
	//order2 := 1<<2
	//fmt.Println(order2)

	//return

	h.Heapify()

	rootIndex, rootValue := h.ExtractRoot()
	rootIndex, rootValue = h.ExtractRoot()
	rootIndex, rootValue = h.ExtractRoot()
	rootIndex, rootValue = h.ExtractRoot()
	rootIndex, rootValue = h.ExtractRoot()

	fmt.Println(fmt.Sprintf("rootIndex=%d, rootValue=%d", rootIndex, rootValue))
	fmt.Println("Hell on earth")
}
