package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

type FirstDepthSearcher struct {
	g                    DirectedGraph
	counterFinishingTime int
	isNodeExplored       []bool
	nodesFinishingTime   []int
	nodesFinishedQueue   []int
	//finishingTimeIndx    int
	//nodesLeader          []int
	//nodeStart            int
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
		}
		if tail < head {
			initialLen = head
		}
		g.edges = make([]DirectedEdge, initialLen)
	}

	// if newLen is required
	newLen := 0
	maxOutOfTailHead := tail
	if tail < head {
		maxOutOfTailHead = head
	}
	if len(g.edges) < maxOutOfTailHead {
		newLen = maxOutOfTailHead
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
	g.edges[tailNodeIdx].heads = append(g.edges[tailNodeIdx].heads, head)

	headNodeIdx := head - 1
	g.edges[headNodeIdx].tail = head
}

func (g *DirectedGraph) checkGraph(grName string) {
	for nodeIdx := 0; nodeIdx < len(g.edges); nodeIdx++ {
		if g.edges[nodeIdx].heads == nil {
			fmt.Println(fmt.Sprintf("For graph %s, at nodeIdx=%d, no node was discovered", grName, nodeIdx))
			continue
		}

		edge := g.edges[nodeIdx]
		if edge.tail != nodeIdx+1 {
			fmt.Println(
				fmt.Sprintf(
					"For graph %s, at nodeIdx=%d, there exists node = %d", grName, nodeIdx, edge.tail))
		}

		if edge.tail == 0 {
			fmt.Println(
				fmt.Sprintf(
					"For graph %s, at nodeIdx=%d, the tail is zero", grName, nodeIdx, edge.tail))
		}
	}
}

func (searcher *FirstDepthSearcher) startSearch(nodeTail int) {
	//fmt.Println("startSearch: Called")
	nodeIdx := nodeTail - 1

	if searcher.isNodeExplored[nodeIdx] {
		//fmt.Println("startSearch: Node", nodeTail, "has been explored. Ignored")
		return
	}
	//fmt.Println("startSearch: Setting nodeTail", nodeTail, "to explored")
	searcher.isNodeExplored[nodeIdx] = true

	if searcher.g.edges[nodeIdx].heads != nil {
		for _, nodeHead := range searcher.g.edges[nodeIdx].heads {
			searcher.startSearch(nodeHead)
		}
	}

	searcher.nodesFinishedQueue = append(searcher.nodesFinishedQueue, nodeTail)
	searcher.nodesFinishingTime[nodeTail-1] = searcher.counterFinishingTime
	searcher.counterFinishingTime++
	//fmt.Println("startSearch: Node", nodeTail, "is fully explored.")
}

func main() {
	g := DirectedGraph{}
	g.PopulateFromFile("./_410e934e6553ac56409b2cb7096a44aa_SCC.txt", false)

	gReversed := DirectedGraph{}
	gReversed.PopulateFromFile("./_410e934e6553ac56409b2cb7096a44aa_SCC.txt", true)

	//g := DirectedGraph{}
	//g.PopulateFromFile("./edges_example_reversed.txt", true)
	//
	//gReversed := DirectedGraph{}
	//gReversed.PopulateFromFile("./edges_example_reversed.txt", false)

	//g.checkGraph("g")
	//gReversed.checkGraph("gReversed")

	firstPassSearcher := FirstDepthSearcher{
		g:                    gReversed,
		counterFinishingTime: 1,
		isNodeExplored:       make([]bool, len(g.edges), len(g.edges)),
		nodesFinishingTime:   make([]int, len(g.edges), len(g.edges)),
		nodesFinishedQueue:   nil,
		//nodesLeader:          nil,
		//nodeStart:            0,
	}

	for nodeTail := len(gReversed.edges) - 1; 0 < nodeTail; nodeTail-- {
		firstPassSearcher.startSearch(nodeTail)
	}

	secondPassSearcher := FirstDepthSearcher{
		g:                    g,
		counterFinishingTime: 1,
		isNodeExplored:       make([]bool, len(g.edges), len(g.edges)),
		nodesFinishingTime:   make([]int, len(g.edges), len(g.edges)),
		nodesFinishedQueue:   nil,
		//nodesLeader:          nil,
		//nodeStart:            0,
	}

	fiveLargestScc := make([]int, 0, 5)
	beforeSearchLen := 0
	afterSearchLen := 0
	for nodeFinishedQueueIdx := len(firstPassSearcher.nodesFinishedQueue) - 1; nodeFinishedQueueIdx >= 0; nodeFinishedQueueIdx-- {
		nodeTail := firstPassSearcher.nodesFinishedQueue[nodeFinishedQueueIdx]

		secondPassSearcher.startSearch(nodeTail)
		afterSearchLen = len(secondPassSearcher.nodesFinishedQueue)

		finishedNodesLength := afterSearchLen - beforeSearchLen

		isStoringFirstFive := len(fiveLargestScc) < 5
		isGreater := false
		if !isStoringFirstFive {
			for _, l := range fiveLargestScc {
				isGreater = finishedNodesLength > l
				if isGreater {
					break
				}
			}
		}

		if isStoringFirstFive {
			fiveLargestScc = append(fiveLargestScc, finishedNodesLength)
			sort.Sort(sort.Reverse(sort.IntSlice(fiveLargestScc)))
		} else if isGreater {
			fiveLargestScc = append(
				[]int{finishedNodesLength},
				fiveLargestScc[0:4]...,
			)

			sort.Sort(sort.Reverse(sort.IntSlice(fiveLargestScc)))
		}
		beforeSearchLen = afterSearchLen
	}

	fmt.Println("main: fiveLargestScc = ", fiveLargestScc)

	fmt.Println("Hell on earth")
}
