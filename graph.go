package graph

import "fmt"

type Graph struct {
	Length       int
	AdjacentList map[string][]string
}

func (g *Graph) AddVertex(node string) {
	if _, ok := g.AdjacentList[node]; !ok {
		g.AdjacentList[node] = []string{}
		g.Length++
	}
}

func (g *Graph) RemoveVertex(node string) {
	if g.Length == 0 {
		return
	}

	if _, ok := g.AdjacentList[node]; ok {
		delete(g.AdjacentList, node)
		g.Length--
	}
}

func (g *Graph) AddEdge(node1, node2 string) {
	g.AdjacentList[node1] = append(g.AdjacentList[node1], node2)
	g.AdjacentList[node2] = append(g.AdjacentList[node2], node1)
}

func (g *Graph) ShowConnections() {
	for i, v := range g.AdjacentList {
		fmt.Println(i, "-->", v)
	}
}

func main() {
	graph := Graph{AdjacentList: map[string][]string{}}
	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")
	graph.AddEdge("3", "1")
	graph.AddEdge("3", "4")
	graph.AddEdge("4", "2")
	graph.AddEdge("4", "5")
	graph.AddEdge("1", "2")
	graph.AddEdge("1", "0")
	graph.AddEdge("0", "2")
	graph.AddEdge("6", "5")
	graph.ShowConnections()
}
