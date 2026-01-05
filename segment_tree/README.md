
## `flood_fill/flood_fill_test.go`

```go
package flood_fill

import "testing"

func TestFloodFill_Core(t *testing.T) {
	t.Parallel()

	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	got := FloodFill(clone2DInt(image), 1, 1, 2)

	want := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}

	assert2DIntEqual(t, want, got)
}

func TestFloodFill_NoOpWhenSameColor(t *testing.T) {
	t.Parallel()

	image := [][]int{
		{0, 0},
		{0, 0},
	}
	got := FloodFill(clone2DInt(image), 0, 0, 0)
	assert2DIntEqual(t, image, got)
}

func TestNumIslands_Core(t *testing.T) {
	t.Parallel()

	grid := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}
	got := NumIslands(clone2DByte(grid))
	if got != 1 {
		t.Fatalf("NumIslands: want 1, got %d", got)
	}

	grid2 := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	got2 := NumIslands(clone2DByte(grid2))
	if got2 != 3 {
		t.Fatalf("NumIslands: want 3, got %d", got2)
	}
}

func TestMaxAreaOfIsland_Core(t *testing.T) {
	t.Parallel()

	grid := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 1, 0, 0},
		{1, 1, 0, 0, 0},
	}
	// Largest island is the plus-shape area = 5
	got := MaxAreaOfIsland(clone2DInt(grid))
	if got != 5 {
		t.Fatalf("MaxAreaOfIsland: want 5, got %d", got)
	}
}

// --- helpers ---

func clone2DInt(a [][]int) [][]int {
	out := make([][]int, len(a))
	for i := range a {
		out[i] = append([]int(nil), a[i]...)
	}
	return out
}

func clone2DByte(a [][]byte) [][]byte {
	out := make([][]byte, len(a))
	for i := range a {
		out[i] = append([]byte(nil), a[i]...)
	}
	return out
}

func assert2DIntEqual(t *testing.T, want, got [][]int) {
	t.Helper()

	if len(want) != len(got) {
		t.Fatalf("row count mismatch: want %d, got %d", len(want), len(got))
	}
	for r := range want {
		if len(want[r]) != len(got[r]) {
			t.Fatalf("col count mismatch row %d: want %d, got %d", r, len(want[r]), len(got[r]))
		}
		for c := range want[r] {
			if want[r][c] != got[r][c] {
				t.Fatalf("cell[%d][%d] mismatch: want %d, got %d\nwant=%v\ngot=%v", r, c, want[r][c], got[r][c], want, got)
			}
		}
	}
}

// TODO (edge cases you should implement yourself):
// - FloodFill: sr/sc out of bounds (define policy).
// - FloodFill: empty image.
// - NumIslands: single cell grids.
// - MaxAreaOfIsland: all zeros / all ones.
// - 8-directional variant (diagonals).
// - Very large grid performance (iterative BFS to avoid recursion overflow).
```
