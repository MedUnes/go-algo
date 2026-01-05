package dfs

import (
	"fmt"
	"testing"
)

func TestDFSRecursive_CoreTraversal(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)

	got := DFSRecursive(g, 1)

	if len(got) != 6 {
		t.Fatalf("DFSRecursive: expected to visit 6 nodes, got %d (%v)", len(got), got)
	}
	if got[0] != 1 {
		t.Fatalf("DFSRecursive: expected first node to be start=1, got %d (%v)", got[0], got)
	}

	// If you rely on deterministic adjacency order, you can enforce expected ordering.
	// Otherwise, validate properties only.
}

func TestDFSIterative_SameAsRecursive(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	rec := DFSRecursive(g, 1)
	iter := DFSIterative(g, 1)

	// This requires both implementations to produce the same order.
	// If your iterative stack pushes neighbors in a different sequence, adjust implementation or relax the test.
	assertIntSliceEqual(t, rec, iter)
}

func TestDFSRecursive_LinearGraph(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	got := DFSRecursive(g, 1)
	want := []int{1, 2, 3, 4}

	assertIntSliceEqual(t, want, got)
}

func TestHasCycle_NoCycle(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(1, 3)

	if HasCycle(g) {
		t.Fatalf("HasCycle: expected false, got true")
	}
}

func TestHasCycle_WithCycle(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	if !HasCycle(g) {
		t.Fatalf("HasCycle: expected true, got false")
	}
}

func TestHasCycle_SelfLoop(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 1)

	if !HasCycle(g) {
		t.Fatalf("HasCycle: expected true for self-loop, got false")
	}
}

func TestFindAllPaths_MultiplePaths(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	got := FindAllPaths(g, 0, 3)

	// Two valid paths; order not specified.
	want := [][]int{{0, 1, 3}, {0, 2, 3}}
	assertPathSetEqual(t, want, got)
}

func TestFindAllPaths_NoPath(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)

	got := FindAllPaths(g, 1, 4)
	if len(got) != 0 {
		t.Fatalf("FindAllPaths: expected 0 paths, got %d (%v)", len(got), got)
	}
}

func TestFindAllPaths_SinglePath(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	got := FindAllPaths(g, 1, 3)
	want := [][]int{{1, 2, 3}}

	assertPathSetEqual(t, want, got)
}

// --- helpers ---

func assertIntSliceEqual(t *testing.T, want, got []int) {
	t.Helper()
	if len(want) != len(got) {
		t.Fatalf("slice length mismatch: want=%d got=%d\nwant=%v\ngot=%v", len(want), len(got), want, got)
	}
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("slice[%d] mismatch: want=%d got=%d\nwant=%v\ngot=%v", i, want[i], got[i], want, got)
		}
	}
}

func assertPathSetEqual(t *testing.T, want, got [][]int) {
	t.Helper()

	toKey := func(p []int) string {
		s := ""
		for i, v := range p {
			if i > 0 {
				s += ","
			}
			s += fmt.Sprintf("%d", v)
		}
		return s
	}

	wm := map[string]int{}
	for _, p := range want {
		wm[toKey(p)]++
	}
	gm := map[string]int{}
	for _, p := range got {
		gm[toKey(p)]++
	}

	if len(wm) != len(gm) {
		t.Fatalf("path set size mismatch:\nwant=%v\ngot=%v", want, got)
	}
	for k, v := range wm {
		if gm[k] != v {
			t.Fatalf("missing path %q (want count=%d got=%d)\nwant=%v\ngot=%v", k, v, gm[k], want, got)
		}
	}
}

// TODO (edge cases you should implement yourself):
// - DFS traversal on cyclic graph should terminate (visited).
// - Disconnected graph: only reachable nodes returned.
// - Deep graph (1000+ depth): iterative DFS avoids recursion overflow.
// - HasCycle for undirected graphs (different logic).
// - FindAllPaths with cycles: define “simple path only” and prevent revisits in current path.
// - Start == target: FindAllPaths returns [[start]].
// - Empty graph / nil graph: define policy.
//
// Bonus TODOs:
// - DFS preorder/postorder explicitly.
// - SCC (Tarjan/Kosaraju).
// - DFS for topological sorting.
