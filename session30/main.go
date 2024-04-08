package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Name                string
	TotalLength         float64
	SourceOfTotalLength *Vertex
	VertexLinks         []*Edge
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

func (g *Graph) AddEdges(vertexIndex int, targets []int, weights []float64) {
	vertex := g.Vertices[vertexIndex]
	vertex.VertexLinks = make([]*Edge, len(targets))
	for i, targetIndex := range targets {
		target := g.Vertices[targetIndex]
		edge := &Edge{Source: vertex, Target: target, Weight: weights[i]}
		vertex.VertexLinks[i] = edge
	}
}

func (g *Graph) Dijkstra() {
	fmt.Println("Dijkstra From Graph Class;")
	for i := 1; i < len(g.Vertices); i++ {
		g.Vertices[i].TotalLength = math.MaxFloat64
	}

	for i := 0; i < len(g.Vertices); i++ {
		currentVertex := g.Vertices[i]
		destinations := currentVertex.VertexLinks
		if destinations == nil {
			continue
		}

		for j := 0; j < len(destinations); j++ {
			currentEdge := destinations[j]
			newLength := currentVertex.TotalLength + currentEdge.Weight
			if newLength < currentEdge.Target.TotalLength {
				currentEdge.Target.TotalLength = newLength
				currentEdge.Target.SourceOfTotalLength = currentVertex
			}
		}
	}

	path := g.Vertices[len(g.Vertices)-1].Name
	v := g.Vertices[len(g.Vertices)-1]
	for v.SourceOfTotalLength != nil {
		path = v.SourceOfTotalLength.Name + " - " + path
		v = v.SourceOfTotalLength
	}
	fmt.Println(g.Vertices[len(g.Vertices)-1].TotalLength)
	fmt.Println(path)

	g.RestoreVertices()
}

func (g *Graph) RestoreVertices() {
	for _, v := range g.Vertices {
		v.TotalLength = 0
		v.SourceOfTotalLength = nil
	}
}

func main() {
	g := NewGraph([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"})

	g.AddEdges(0, []int{1, 2, 3}, []float64{2, 4, 3})
	g.AddEdges(1, []int{4, 5, 6}, []float64{7, 4, 6})
	g.AddEdges(2, []int{4, 5, 6}, []float64{3, 2, 4})
	g.AddEdges(3, []int{4, 5, 6}, []float64{4, 1, 5})
	g.AddEdges(4, []int{7, 8}, []float64{1, 4})
	g.AddEdges(5, []int{7, 8}, []float64{6, 3})
	g.AddEdges(6, []int{7, 8}, []float64{3, 3})
	g.AddEdges(7, []int{9}, []float64{3})
	g.AddEdges(8, []int{9}, []float64{4})

	g.Dijkstra()
}
