package topological_sort

import (
	"testing"
)

func TestTopologicalSortKahn_ValidOrders(t *testing.T) {
	t.Parallel()

	g := NewDirectedGraph(6)
	// 5->2, 5->0, 4->0, 4->1, 2->3, 3->1
	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	order := TopologicalSortKahn(g)

	assertTopoOrder(t, g.NumVertices(), edgesFromGraph(t, g), order)
}

func TestTopologicalSortDFS_ValidOrders(t *testing.T) {
	t.Parallel()

	g := NewDirectedGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	order := TopologicalSortDFS(g)
	assertTopoOrder(t, g.NumVertices(), edgesFromGraph(t, g), order)
}

func TestTopologicalSort_CycleBehavior(t *testing.T) {
	t.Parallel()

	// Define expected behavior:
	// - For Kahn: return empty slice on cycle (recommended for interviews), OR
	// - return partial order (then tests should change).
	// Pick ONE policy and keep it consistent.
	g := NewDirectedGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)

	orderK := TopologicalSortKahn(g)
	if len(orderK) != 0 {
		t.Fatalf("on cycle: expected empty order, got %v", orderK)
	}

	// DFS version: also expected empty on cycle for this kata.
	orderD := TopologicalSortDFS(g)
	if len(orderD) != 0 {
		t.Fatalf("on cycle: expected empty order, got %v", orderD)
	}
}

func TestCanFinishCourses_Core(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		n         int
		prereqs   [][]int
		canFinish bool
	}{
		{
			"simple_possible",
			2,
			[][]int{{1, 0}}, // 0 -> 1
			true,
		},
		{
			"simple_cycle",
			2,
			[][]int{{1, 0}, {0, 1}},
			false,
		},
		{
			"disconnected_ok",
			4,
			[][]int{{2, 1}},
			true,
		},
		{
			"long_chain_ok",
			4,
			[][]int{{1, 0}, {2, 1}, {3, 2}},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := CanFinishCourses(tt.n, tt.prereqs)
			if got != tt.canFinish {
				t.Fatalf("CanFinishCourses(n=%d, prereqs=%v): want %v, got %v", tt.n, tt.prereqs, tt.canFinish, got)
			}
		})
	}
}

// --- test helpers ---

func edgesFromGraph(t *testing.T, g *DirectedGraph) [][2]int {
	// If you don't want Edges() in your API, replace this by exposing adjacency
	// or re-constructing edges in tests.
	edges := g.Edges()
	if edges == nil {
		t.Fatalf("DirectedGraph.Edges() returned nil; expose edges for tests or adjust test helper")
	}
	return edges
}

func assertTopoOrder(t *testing.T, n int, edges [][2]int, order []int) {
	t.Helper()

	if len(order) != n {
		t.Fatalf("topo order must include all vertices: want len=%d, got %d (%v)", n, len(order), order)
	}

	pos := make([]int, n)
	seen := make([]bool, n)
	for i, v := range order {
		if v < 0 || v >= n {
			t.Fatalf("vertex out of range: %d", v)
		}
		if seen[v] {
			t.Fatalf("duplicate vertex in topo order: %d in %v", v, order)
		}
		seen[v] = true
		pos[v] = i
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		if pos[u] >= pos[v] {
			t.Fatalf("invalid topo order: edge %d->%d violated by order %v", u, v, order)
		}
	}
}

// TODO (edge cases you should implement yourself):
// - Graph with isolated vertices only.
// - Multiple components with edges in each.
// - Large DAG performance (n=100k, m=200k) smoke test.
// - Deterministic output requirement (e.g., lexicographically smallest) using heap.
// - Invalid edges (u or v out of range): decide panic vs ignore vs error (prefer error).
