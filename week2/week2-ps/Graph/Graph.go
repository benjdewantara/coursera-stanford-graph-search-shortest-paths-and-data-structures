package Graph

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type WeightedEdge struct {
	Tail    int
	Heads   []int
	Weights []int
}

type Graph struct {
	Edges []WeightedEdge
}

func (g *Graph) PopulateFromFile(filepath string) {
	contentBytes, _ := ioutil.ReadFile(filepath)
	for _, edgeStr := range strings.Split(string(contentBytes), "\n") {
		edgeStr = strings.TrimSuffix(edgeStr, " ")
		edgeStrSplit := strings.Split(edgeStr, "\t")
		tail, _ := strconv.Atoi(edgeStrSplit[0])
		e := WeightedEdge{
			Tail: tail,
		}
		for _, headWeightStr := range edgeStrSplit[1:] {
			headWeightStrSplit := strings.Split(headWeightStr, ",")
			if len(headWeightStrSplit) != 2 {
				continue
			}
			head, _ := strconv.Atoi(headWeightStrSplit[0])
			weight, _ := strconv.Atoi(headWeightStrSplit[1])
			e.Heads = append(e.Heads, head)
			e.Weights = append(e.Weights, weight)
		}

		g.Edges = append(g.Edges, e)
	}
}
