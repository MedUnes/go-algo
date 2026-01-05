# Two Heaps Pattern

## 0) Interview framing (Google/FAANG)
**What they test:** priority queue mechanics, maintaining invariants, and median logic under streaming input.  
**What you should say out loud:**
- "Left heap is a max-heap (lower half), right heap is a min-heap (upper half).”
- "Invariant: size(left) == size(right) or size(left) == size(right)+1.”
- "Median is top of left (odd) or avg of both tops (even).”

---

## 1) Overview
Two Heaps maintains two balanced heaps to support fast median retrieval or access to both extremes in a dynamic dataset.

## 2) Problem(s) it solves
- Median of a stream
- Sliding window median (with deletions)
- Meeting rooms / scheduling with end-time heap (related use)

## 3) Core idea & invariants
### Invariants
- All elements in `maxHeap` (left) ≤ all elements in `minHeap` (right).
- `len(maxHeap) == len(minHeap)` or `len(maxHeap) == len(minHeap)+1`.
- Median computed from roots based on parity.

### Why it works
By keeping halves separated and balanced, median is always at the boundary.

## 4) Typical formulations
- "Design a data structure that supports addNum and findMedian.”
- "Return running medians.”

## 5) Complexity
- **Insert:** O(log n)
- **FindMedian:** O(1)
- **Space:** O(n)

## 6) Go notes
- Use `container/heap`. Implement `heap.Interface` with pointer receivers.
- Max-heap can be implemented by inverting `Less`.

## 7) Pitfalls
1. Not rebalancing after insertion.
2. Violating ordering invariant between heaps.
3. Overflow when averaging two ints (use float64 with care).

## 8) Practice katas
- Find Median from Data Stream (LC #295)
- Sliding Window Median (LC #480)
- Meeting Rooms II (LC #253)

## 9) Variations
- Sliding window with lazy deletions
- Kth largest (single min-heap)
- Three-heaps for quantiles

## 10) Further reading
- https://pkg.go.dev/container/heap
- https://leetcode.com/problems/find-median-from-data-stream/

---

## Contract (your Go API)

```go
package two_heaps

type MedianFinder struct {
	// your fields
}

func NewMedianFinder() *MedianFinder
func (mf *MedianFinder) AddNum(num int)
func (mf *MedianFinder) FindMedian() float64

type MaxHeap []int
type MinHeap []int
```