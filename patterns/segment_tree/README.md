# Segment Tree

## 0) Interview framing (Google/FAANG)
**What they test:** range queries + point updates in O(log n), clean indexing, and correctness of recursion.  
**What you should say out loud:**
- “Segment tree stores aggregates over intervals.”
- “Build O(n), query O(log n), update O(log n).”
- “Each node covers a segment; children split it in half.”

---

## 1) Overview
A Segment Tree supports fast range queries (sum/min/max) while allowing updates, ideal when you have many queries and modifications.

## 2) Problems it solves
- Range Sum Query – Mutable (LC #307)
- Range Minimum Query
- General interval aggregates

## 3) Core invariants
- Node value equals aggregate over its segment.
- Update fixes path from leaf to root.
- Query splits into disjoint covered segments.

## 4) Complexity
- Build: O(n)
- Query: O(log n)
- Update: O(log n)

## 5) Go notes
- Use iterative or recursive; recursive is easier to reason about in interviews.
- Validate boundaries clearly (left/right inclusive).
- Decide behavior for empty input.

## 6) Pitfalls
1. Off-by-one segment boundaries.
2. Wrong neutral element for min/sum.
3. Not handling updates correctly.

## 7) Practice katas
- Range Sum Query - Mutable (LC #307)
- Count of Range Sum (LC #327) (advanced)

## 8) Further reading
- https://cp-algorithms.com/data_structures/segment_tree.html
- https://leetcode.com/problems/range-sum-query-mutable/

---

## Contract (your Go API)

```go
package segment_tree

type SegmentTree struct {
	tree []int
	n    int
}

func NewSegmentTree(nums []int) *SegmentTree
func (st *SegmentTree) Update(index int, val int)
func (st *SegmentTree) RangeSum(left int, right int) int
func (st *SegmentTree) RangeMin(left int, right int) int
```