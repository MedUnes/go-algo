package bfs

type Graph struct {
	// your representation
}

func NewGraph() *Graph
func (g *Graph) AddEdge(u, v int)
func (g *Graph) AddUndirectedEdge(u, v int)

func BFS(g *Graph, start int) []int
func BFSWithLevels(g *Graph, start int) [][]int

// ShortestPath returns (path, distance). If no path: (nil, -1).
func ShortestPath(g *Graph, start, target int) ([]int, int)
