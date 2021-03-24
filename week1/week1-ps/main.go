package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type DirectedEdge struct {
	tail  int
	heads []int
}

type DirectedGraph struct {
	edges []DirectedEdge
}

func (g *DirectedGraph) PopulateFromFile(filepath string, isHeadTailReversed bool) {
	contentBytes, _ := ioutil.ReadFile(filepath)
	for _, edgeStr := range strings.Split(string(contentBytes), "\n") {
		edgeStr = strings.TrimSuffix(edgeStr, " ")
		edgeStrSplit := strings.Split(edgeStr, " ")
		tail, _ := strconv.Atoi(edgeStrSplit[0])
		head, _ := strconv.Atoi(edgeStrSplit[1])

		if isHeadTailReversed {
			g.AddDirectedEdge(head, tail)
		} else {
			g.AddDirectedEdge(tail, head)
		}
	}
}

func (g *DirectedGraph) AddDirectedEdge(tail int, head int) {
	// if g.edges is still nil
	if g.edges == nil {
		initialLen := 8
		if initialLen < tail {
			initialLen = tail
		} else if initialLen < head {
			initialLen = head
		}
		g.edges = make([]DirectedEdge, initialLen)
	}

	// if newLen is required
	newLen := 0
	if len(g.edges) < int(tail) {
		newLen = tail
	} else if len(g.edges) < int(head) {
		newLen = head
	}

	// grow the capacity of g.edges
	if newLen != 0 {
		if int(newLen) < cap(g.edges) {
			g.edges = g.edges[0:newLen]
		} else {
			newEdges := make([]DirectedEdge, newLen+1, int((newLen*5+4)/4))
			copy(newEdges, g.edges)
			g.edges = newEdges
		}
	}

	tailNodeIdx := tail - 1
	g.edges[tailNodeIdx].tail = tail
	if g.edges[tailNodeIdx].heads == nil {
		g.edges[tailNodeIdx].heads = make([]int, 0, 8)
	}

	// grow the capacity of array of head nodes pointed to by tail node
	headLen := len(g.edges[tailNodeIdx].heads)
	if cap(g.edges[tailNodeIdx].heads) <= headLen {
		newHeadNodes := make([]int, headLen, int(headLen*5/4))
		copy(newHeadNodes, g.edges[tailNodeIdx].heads)
		g.edges[tailNodeIdx].heads = newHeadNodes
	}

	g.edges[tailNodeIdx].heads = g.edges[tailNodeIdx].heads[0 : headLen+1]
	g.edges[tailNodeIdx].heads[headLen] = head
}

func main() {
	g := DirectedGraph{}
	g.PopulateFromFile("./_410e934e6553ac56409b2cb7096a44aa_SCC.txt", false)

	g_reversed := DirectedGraph{}
	g_reversed.PopulateFromFile("./_410e934e6553ac56409b2cb7096a44aa_SCC.txt", true)

	for nodeIdx := 0; nodeIdx < len(g.edges); nodeIdx++ {
		if g.edges[nodeIdx].heads == nil {
			fmt.Println(fmt.Sprintf("For graph g, at nodeIdx=%d, no node was discovered", nodeIdx))
			continue
		}

		edge := g.edges[nodeIdx]
		if edge.tail != nodeIdx+1 {
			fmt.Println(fmt.Sprintf("For graph g, at nodeIdx=%d, there exists node = %d", nodeIdx, edge.tail))
		}
	}

	for nodeIdx := 0; nodeIdx < len(g_reversed.edges); nodeIdx++ {
		if g_reversed.edges[nodeIdx].heads == nil {
			fmt.Println(fmt.Sprintf("For graph g_reversed, at nodeIdx=%d, no node was discovered", nodeIdx))
			continue
		}

		edge := g_reversed.edges[nodeIdx]
		if edge.tail != nodeIdx+1 {
			fmt.Println(
				fmt.Sprintf(
					"For graph g_reversed, at nodeIdx=%d, there exists node = %d", nodeIdx, edge.tail))
		}
	}

	fmt.Println("Hell on earth")
}
