package topological_sort

type DirectedGraph struct {
	// your representation
}

func NewDirectedGraph(n int) *DirectedGraph
func (g *DirectedGraph) AddEdge(u, v int)

func (g *DirectedGraph) NumVertices() int
func (g *DirectedGraph) Edges() [][2]int // optional helper for tests; you can expose differently

func TopologicalSortKahn(g *DirectedGraph) []int
func TopologicalSortDFS(g *DirectedGraph) []int
func CanFinishCourses(numCourses int, prerequisites [][]int) bool
