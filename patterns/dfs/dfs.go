package dfs

type Graph struct {
	// your representation
}

func NewGraph() *Graph
func (g *Graph) AddEdge(u, v int)

func DFSRecursive(g *Graph, start int) []int
func DFSIterative(g *Graph, start int) []int

// HasCycle detects a cycle in a directed graph.
func HasCycle(g *Graph) bool

// FindAllPaths returns all simple paths from start to target.
// If you allow cycles, define how you prevent infinite paths (usually disallow revisiting nodes in current path).
func FindAllPaths(g *Graph, start, target int) [][]int
