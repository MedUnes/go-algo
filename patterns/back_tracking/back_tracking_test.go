package backtracking

import (
	"fmt"
	"testing"
)

func TestPermutations_Core(t *testing.T) {
	t.Parallel()

	in := []int{1, 2, 3}
	got := Permutations(in)

	// Expect 3! = 6 permutations (order not specified).
	if len(got) != 6 {
		t.Fatalf("Permutations(%v): expected 6 permutations, got %d (%v)", in, len(got), got)
	}

	// Validate: all permutations are unique and contain the same multiset.
	assertPermutationsOf(t, in, got)
}

func TestPermutations_Single(t *testing.T) {
	t.Parallel()

	in := []int{42}
	got := Permutations(in)
	if len(got) != 1 || len(got[0]) != 1 || got[0][0] != 42 {
		t.Fatalf("Permutations(%v): expected [[42]], got %v", in, got)
	}
}

func TestCombinations_Core(t *testing.T) {
	t.Parallel()

	got := Combinations(4, 2)

	// C(4,2)=6
	if len(got) != 6 {
		t.Fatalf("Combinations(4,2): expected 6, got %d (%v)", len(got), got)
	}

	// Validate each combination:
	// - length k
	// - strictly increasing
	// - values in [1..n]
	for _, comb := range got {
		if len(comb) != 2 {
			t.Fatalf("invalid combination length: %v", comb)
		}
		if comb[0] >= comb[1] {
			t.Fatalf("combination not increasing: %v", comb)
		}
		if comb[0] < 1 || comb[1] > 4 {
			t.Fatalf("combination out of range: %v", comb)
		}
	}

	// Validate uniqueness (order not specified).
	assertUniqueSlices(t, got)
}

func TestSolveSudoku_Core(t *testing.T) {
	t.Parallel()

	// Classic solvable Sudoku ('.' means empty)
	board := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}

	ok := SolveSudoku(board)
	if !ok {
		t.Fatalf("SolveSudoku: expected true, got false")
	}

	assertSudokuSolved(t, board)
}

func TestNQueens_Core(t *testing.T) {
	t.Parallel()

	solutions := NQueens(4)
	// N=4 has 2 solutions.
	if len(solutions) != 2 {
		t.Fatalf("NQueens(4): expected 2 solutions, got %d", len(solutions))
	}
	for _, sol := range solutions {
		assertValidNQueensSolution(t, 4, sol)
	}
}

// --- helpers ---

func assertPermutationsOf(t *testing.T, in []int, perms [][]int) {
	t.Helper()

	wantCounts := toCountMap(in)

	seen := map[string]struct{}{}
	for _, p := range perms {
		if len(p) != len(in) {
			t.Fatalf("permutation has wrong length: want %d, got %d (%v)", len(in), len(p), p)
		}
		if !equalCountMap(wantCounts, toCountMap(p)) {
			t.Fatalf("permutation has wrong elements: in=%v perm=%v", in, p)
		}
		k := keyOf(p)
		if _, ok := seen[k]; ok {
			t.Fatalf("duplicate permutation found: %v", p)
		}
		seen[k] = struct{}{}
	}
}

func assertUniqueSlices(t *testing.T, xs [][]int) {
	t.Helper()
	seen := map[string]struct{}{}
	for _, x := range xs {
		k := keyOf(x)
		if _, ok := seen[k]; ok {
			t.Fatalf("duplicate slice found: %v", x)
		}
		seen[k] = struct{}{}
	}
}

func assertSudokuSolved(t *testing.T, b [][]byte) {
	t.Helper()

	// Validate rows/cols/3x3 each has digits 1..9 exactly once.
	for i := 0; i < 9; i++ {
		row := make([]byte, 0, 9)
		col := make([]byte, 0, 9)
		for j := 0; j < 9; j++ {
			row = append(row, b[i][j])
			col = append(col, b[j][i])
		}
		assertDigits1to9(t, row, fmt.Sprintf("row %d", i))
		assertDigits1to9(t, col, fmt.Sprintf("col %d", i))
	}
	for br := 0; br < 3; br++ {
		for bc := 0; bc < 3; bc++ {
			box := make([]byte, 0, 9)
			for r := br * 3; r < br*3+3; r++ {
				for c := bc * 3; c < bc*3+3; c++ {
					box = append(box, b[r][c])
				}
			}
			assertDigits1to9(t, box, fmt.Sprintf("box %d,%d", br, bc))
		}
	}
}

func assertDigits1to9(t *testing.T, xs []byte, label string) {
	t.Helper()
	seen := make([]bool, 10)
	for _, ch := range xs {
		if ch < '1' || ch > '9' {
			t.Fatalf("%s: invalid char %q in %q", label, ch, xs)
		}
		d := int(ch - '0')
		if seen[d] {
			t.Fatalf("%s: duplicate digit %d in %q", label, d, xs)
		}
		seen[d] = true
	}
}

func assertValidNQueensSolution(t *testing.T, n int, board []string) {
	t.Helper()

	if len(board) != n {
		t.Fatalf("NQueens solution: expected %d rows, got %d", n, len(board))
	}

	cols := make([]bool, n)
	diag1 := make(map[int]bool) // r-c
	diag2 := make(map[int]bool) // r+c
	queens := 0

	for r := 0; r < n; r++ {
		if len(board[r]) != n {
			t.Fatalf("row %d: expected length %d, got %d", r, n, len(board[r]))
		}
		rowQueens := 0
		for c := 0; c < n; c++ {
			ch := board[r][c]
			if ch == 'Q' {
				rowQueens++
				queens++
				if cols[c] {
					t.Fatalf("duplicate queen in col %d", c)
				}
				cols[c] = true
				if diag1[r-c] {
					t.Fatalf("duplicate queen on diag r-c=%d", r-c)
				}
				diag1[r-c] = true
				if diag2[r+c] {
					t.Fatalf("duplicate queen on diag r+c=%d", r+c)
				}
				diag2[r+c] = true
			} else if ch != '.' {
				t.Fatalf("invalid char %q (expected 'Q' or '.')", ch)
			}
		}
		if rowQueens != 1 {
			t.Fatalf("row %d: expected exactly 1 queen, got %d", r, rowQueens)
		}
	}
	if queens != n {
		t.Fatalf("expected %d queens, got %d", n, queens)
	}
}

func toCountMap(xs []int) map[int]int {
	m := map[int]int{}
	for _, v := range xs {
		m[v]++
	}
	return m
}

func equalCountMap(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, av := range a {
		if bv, ok := b[k]; !ok || bv != av {
			return false
		}
	}
	return true
}

func keyOf(xs []int) string {
	// stable representation to compare slices ignoring ordering of results
	// (does NOT sort; preserves order within the slice)
	s := ""
	for i, v := range xs {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprintf("%d", v)
	}
	return s
}

// TODO (edge cases you should implement yourself):
// - Permutations with duplicates in input: decide if you return duplicates or unique permutations.
// - Combinations: k==0, k>n, n==0 (define expected behavior).
// - SolveSudoku: unsolvable board should return false and not corrupt board.
// - NQueens: n==1, n==2, n==3 edge results.
// - Performance: NQueens(12) smoke test with pruning.
//
// Bonus TODOs:
// - Subsets (LC #78)
// - Word Search (LC #79)
// - Combination Sum variants
