package bfs

import (
	"testing"
)

func TestBFS_CoreTraversal(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	// 1 -> {2,3}, 2 -> {4,5}, 3 -> {6}
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)

	got := BFS(g, 1)

	// This test assumes deterministic neighbor iteration order (because you used AddEdge in a fixed order).
	// If your graph uses map iteration, this will be nondeterministic; either:
	// - store adjacency as slices, or
	// - sort neighbors during traversal, or
	// - change this test to validate properties instead of exact order.
	want := []int{1, 2, 3, 4, 5, 6}

	assertIntSliceEqual(t, want, got)
}

func TestBFS_SingleNode(t *testing.T) {
	t.Parallel()

	g := NewGraph()

	got := BFS(g, 1)
	want := []int{1}

	assertIntSliceEqual(t, want, got)
}

func TestBFS_LinearGraph(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	got := BFS(g, 1)
	want := []int{1, 2, 3, 4}

	assertIntSliceEqual(t, want, got)
}

func TestBFSWithLevels_Tree(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	// Perfect-ish binary tree from 1
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	got := BFSWithLevels(g, 1)
	want := [][]int{
		{1},
		{2, 3},
		{4, 5, 6, 7},
	}

	assert2DIntSliceEqual(t, want, got)
}

func TestShortestPath_Found(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	// Undirected:
	// 1-2-4
	//  \3/
	g.AddUndirectedEdge(1, 2)
	g.AddUndirectedEdge(1, 3)
	g.AddUndirectedEdge(2, 4)
	g.AddUndirectedEdge(3, 4)

	path, dist := ShortestPath(g, 1, 4)
	if dist != 2 {
		t.Fatalf("ShortestPath distance: want 2, got %d (path=%v)", dist, path)
	}

	// Multiple valid answers: [1,2,4] OR [1,3,4].
	// Validate properties instead of exact path.
	if len(path) != 3 || path[0] != 1 || path[2] != 4 {
		t.Fatalf("ShortestPath path shape invalid: got %v", path)
	}
}

func TestShortestPath_NoPath(t *testing.T) {
	t.Parallel()

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)

	path, dist := ShortestPath(g, 1, 4)
	if path != nil || dist != -1 {
		t.Fatalf("ShortestPath no-path: want (nil,-1), got (path=%v, dist=%d)", path, dist)
	}
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

func assert2DIntSliceEqual(t *testing.T, want, got [][]int) {
	t.Helper()
	if len(want) != len(got) {
		t.Fatalf("2D slice row mismatch: want=%d got=%d\nwant=%v\ngot=%v", len(want), len(got), want, got)
	}
	for r := range want {
		assertIntSliceEqual(t, want[r], got[r])
	}
}

// TODO (edge cases you should implement yourself):
// - Graph with cycles should terminate (mark visited on enqueue).
// - Disconnected graph: only reachable nodes returned.
// - Start node == target in ShortestPath: distance 0 and path [start].
// - Nil graph / empty adjacency: define policy (panic vs empty vs single).
// - Self-loops.
// - Very large graph (10k+ nodes) performance smoke test.
//
// Bonus TODOs (variations):
// - BFS on 2D grid (maze).
// - Multi-source BFS.
// - Bidirectional BFS.
// - BFS with early termination when target found.
// - All nodes at exact distance K.
