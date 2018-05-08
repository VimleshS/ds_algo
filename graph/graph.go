package graph

import (
	"fmt"
)

type Queue struct {
	_q []int
}

func (q *Queue) Push(v int) {
	q._q = append(q._q, v)
}

func (q *Queue) Pop() int {
	_t := q._q[0]
	q._q = q._q[1:]
	return _t
}

func (q *Queue) Len() int {
	return len(q._q)
}

type AdjacencyMatrices struct {
	NumVertices int
	Directed    bool
	Matrices    *([][]int)
}

func NewAdjacencyMatrices(numVertices int, directed bool) *AdjacencyMatrices {
	a := make([][]int, numVertices)
	for i := range a {
		a[i] = make([]int, numVertices)
	}
	return &AdjacencyMatrices{NumVertices: numVertices, Matrices: &a, Directed: directed}
}

func (g *AdjacencyMatrices) PrintMatrices() {
	for i := 0; i < g.NumVertices; i++ {
		fmt.Println((*g.Matrices)[i])
	}
}

func (g *AdjacencyMatrices) GetEdgeWeight(v1, v2 int) int {
	return (*g.Matrices)[v1][v2]
}

func (g *AdjacencyMatrices) GetAdjacents(v1 int) []int {
	adjacents := []int{}
	for i := 0; i < g.NumVertices; i++ {
		if (*g.Matrices)[v1][i] > 0 {
			adjacents = append(adjacents, i)
		}
	}
	return adjacents
}

func (g *AdjacencyMatrices) GetIndegree(v1 int) int {
	indegree := 0
	for i := 0; i < g.NumVertices; i++ {
		if (*g.Matrices)[i][v1] > 0 {
			indegree += (*g.Matrices)[i][v1]
		}
	}
	return indegree
}

func (g *AdjacencyMatrices) AddEdge(v1, v2, value int) error {
	if v1 > g.NumVertices || v2 > g.NumVertices {
		return fmt.Errorf("Invalid indices for vertexs")
	}
	(*g.Matrices)[v1][v2] = value
	if !g.Directed {
		(*g.Matrices)[v2][v1] = value
	}
	return nil
}

func TestAdjacencyMatrices(g *AdjacencyMatrices) {
	for i := 0; i < g.NumVertices; i++ {
		fmt.Printf("Adjancent to %d %v \n", i, g.GetAdjacents(i))
	}

	for i := 0; i < g.NumVertices; i++ {
		fmt.Printf("Indegree %d %d \n", i, g.GetIndegree(i))
	}

	for i := 0; i < g.NumVertices; i++ {
		for _, j := range g.GetAdjacents(i) {
			fmt.Printf("Edge Weight %d %s %d %s %d \n", i, " ", j, " ", g.GetEdgeWeight(i, j))
		}

	}
	g.PrintMatrices()
}

func TopologicalSort() {
	g := NewAdjacencyMatrices(9, true)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 7, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 7, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(5, 6, 1)
	g.AddEdge(3, 6, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(6, 8, 1)

	// TestAdjacencyMatrices(g)

	q := Queue{}
	indegree := map[int]int{}
	for i := 0; i < g.NumVertices; i++ {
		indegree[i] = g.GetIndegree(i)
		if indegree[i] == 0 {
			q.Push(i)
		}
	}

	topoSorted := []int{}
	for q.Len() > 0 {
		v := q.Pop()
		topoSorted = append(topoSorted, v)

		adjs := g.GetAdjacents(v)
		for _, adj := range adjs {
			indegree[adj] = indegree[adj] - 1
			if indegree[adj] == 0 {
				q.Push(adj)
			}
		}
	}
	fmt.Println(topoSorted)
}
