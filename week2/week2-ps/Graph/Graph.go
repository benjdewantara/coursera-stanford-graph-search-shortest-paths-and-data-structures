package Graph

import (
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
	Edges []WeightedEdge
}

func (g *Graph) PopulateFromFile(filepath string) {
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

		g.Edges = append(g.Edges, e)
	}
}
