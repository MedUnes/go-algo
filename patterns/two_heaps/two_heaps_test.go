package two_heaps

import (
	"container/heap"
	"testing"
)

func TestMedianFinder_Core(t *testing.T) {
	t.Parallel()

	mf := NewMedianFinder()
	mf.AddNum(5)
	if got := mf.FindMedian(); got != 5.0 {
		t.Fatalf("median after [5]: want 5.0, got %f", got)
	}

	mf.AddNum(15)
	if got := mf.FindMedian(); got != 10.0 {
		t.Fatalf("median after [5,15]: want 10.0, got %f", got)
	}

	mf.AddNum(1)
	if got := mf.FindMedian(); got != 5.0 {
		t.Fatalf("median after [5,15,1]: want 5.0, got %f", got)
	}

	mf.AddNum(3)
	if got := mf.FindMedian(); got != 4.0 {
		t.Fatalf("median after [5,15,1,3]: want 4.0, got %f", got)
	}
}

func TestMedianFinder_OddEven(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want float64
	}{
		{"odd", []int{5, 15, 1, 3, 8}, 5.0},
		{"even", []int{1, 2, 3, 4}, 2.5},
		{"dups", []int{1, 1, 1, 2, 2, 2}, 1.5},
		{"negatives", []int{-1, -2, -3, -4, -5}, -3.0},
		{"mixed", []int{-10, 5, 0, -5, 10}, 0.0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mf := NewMedianFinder()
			for _, x := range tt.nums {
				mf.AddNum(x)
			}
			if got := mf.FindMedian(); got != tt.want {
				t.Fatalf("FindMedian(%v): want %f, got %f", tt.nums, tt.want, got)
			}
		})
	}
}

func TestMaxHeap_BasicOperations(t *testing.T) {
	t.Parallel()

	h := &MaxHeap{}
	heap.Init(h)

	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 5)
	heap.Push(h, 2)

	// max is 5 at index 0
	if (*h)[0] != 5 {
		t.Fatalf("MaxHeap top: want 5, got %d", (*h)[0])
	}
	if got := heap.Pop(h).(int); got != 5 {
		t.Fatalf("MaxHeap pop: want 5, got %d", got)
	}
	if (*h)[0] != 3 {
		t.Fatalf("MaxHeap new top: want 3, got %d", (*h)[0])
	}
}

func TestMinHeap_BasicOperations(t *testing.T) {
	t.Parallel()

	h := &MinHeap{}
	heap.Init(h)

	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 5)
	heap.Push(h, 2)

	// min is 1 at index 0
	if (*h)[0] != 1 {
		t.Fatalf("MinHeap top: want 1, got %d", (*h)[0])
	}
	if got := heap.Pop(h).(int); got != 1 {
		t.Fatalf("MinHeap pop: want 1, got %d", got)
	}
	if (*h)[0] != 2 {
		t.Fatalf("MinHeap new top: want 2, got %d", (*h)[0])
	}
}

// TODO (edge cases you should implement yourself):
// - Empty MedianFinder: define behavior (panic vs 0 vs error).
// - Very large ints near MaxInt: avoid overflow when averaging.
// - Randomized property test: compare to sorted slice median.
// - Performance benchmark: 1e6 inserts.
// - Alternating high/low values: stress rebalancing.
//
// Bonus TODOs:
// - Sliding window median with lazy deletions.
// - Generic heap wrappers (Go 1.18+ generics).
