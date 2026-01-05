package segment_tree

import "testing"

func TestSegmentTree_RangeSum_Core(t *testing.T) {
	t.Parallel()

	st := NewSegmentTree([]int{1, 3, 5})
	if got := st.RangeSum(0, 2); got != 9 {
		t.Fatalf("RangeSum(0,2): want 9, got %d", got)
	}

	st.Update(1, 2) // [1,2,5]
	if got := st.RangeSum(0, 2); got != 8 {
		t.Fatalf("after Update(1,2) RangeSum(0,2): want 8, got %d", got)
	}
	if got := st.RangeSum(1, 2); got != 7 {
		t.Fatalf("RangeSum(1,2): want 7, got %d", got)
	}
}

func TestSegmentTree_RangeMin_Core(t *testing.T) {
	t.Parallel()

	st := NewSegmentTree([]int{4, 2, 7, 1, 3})
	if got := st.RangeMin(0, 4); got != 1 {
		t.Fatalf("RangeMin(0,4): want 1, got %d", got)
	}
	if got := st.RangeMin(0, 1); got != 2 {
		t.Fatalf("RangeMin(0,1): want 2, got %d", got)
	}

	st.Update(3, 10) // [4,2,7,10,3]
	if got := st.RangeMin(0, 4); got != 2 {
		t.Fatalf("after Update(3,10) RangeMin(0,4): want 2, got %d", got)
	}
}

func TestSegmentTree_SingleElement(t *testing.T) {
	t.Parallel()

	st := NewSegmentTree([]int{42})
	if got := st.RangeSum(0, 0); got != 42 {
		t.Fatalf("RangeSum(0,0): want 42, got %d", got)
	}
	if got := st.RangeMin(0, 0); got != 42 {
		t.Fatalf("RangeMin(0,0): want 42, got %d", got)
	}

	st.Update(0, -1)
	if got := st.RangeSum(0, 0); got != -1 {
		t.Fatalf("after update RangeSum(0,0): want -1, got %d", got)
	}
}

// TODO (edge cases you should implement yourself):
// - Empty input: NewSegmentTree(nil) behavior (panic vs usable zero-value).
// - Out-of-range queries/updates: return 0/error/panic (define policy).
// - RangeSum(left>right): define policy.
// - Large random array property test (compare against prefix sums for sums).
// - Min neutral element correctness (e.g., use +Inf).
// - Performance benchmark (n=200k, q=200k).
//
// Bonus TODOs:
// - Lazy propagation for range updates.
// - Generic segment tree with custom combine function.
