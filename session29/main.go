package main

import (
	"fmt"
)

type Vertex struct {
	Name    string
	Visited bool
	Edges   []*Edge
}

type Edge struct {
	Weight float64
	Source *Vertex
	Target *Vertex
}

type Graph struct {
	Vertices []*Vertex
}

func NewGraph(names []string) *Graph {
	graph := &Graph{}
	for _, name := range names {
		graph.Vertices = append(graph.Vertices, &Vertex{Name: name})
	}
	return graph
}

func (g *Graph) AddEdges(vertexIndex int, targets []int) {
	vertex := g.Vertices[vertexIndex]
	for _, targetIndex := range targets {
		target := g.Vertices[targetIndex]
		edge := &Edge{Source: vertex, Target: target}
		vertex.Edges = append(vertex.Edges, edge)
	}
}

func (g *Graph) ResetVertices() {
	for _, v := range g.Vertices {
		v.Visited = false
	}
}

func (g *Graph) DFS() {
	fmt.Println("DFS From Graph Class:")
	g.DFSRecursion(g.Vertices[0])
	g.ResetVertices()

}

func (g *Graph) DFSRecursion(currentVertex *Vertex) {
	currentVertex.Visited = true
	for _, edge := range currentVertex.Edges {
		if !edge.Target.Visited {
			fmt.Printf("%s - %s\n", currentVertex.Name, edge.Target.Name)
			g.DFSRecursion(edge.Target)
		}
	}

}

func main() {
	g := NewGraph([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"})
	g.AddEdges(0, []int{1, 2})
	g.AddEdges(1, []int{0, 3, 4})
	g.AddEdges(2, []int{0, 3, 5})
	g.AddEdges(3, []int{1, 2, 4})
	g.AddEdges(4, []int{1, 5})
	g.AddEdges(5, []int{2, 3, 4, 7})
	g.AddEdges(6, []int{7, 8})
	g.AddEdges(7, []int{5, 6, 8})
	g.AddEdges(8, []int{6, 7})

	g.DFS()
}
