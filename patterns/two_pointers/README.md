# Two Pointers Pattern

## 0) Interview framing (Google/FAANG)
**What they test:** invariants with indices, linear-time scanning, and correct pointer movement.  
**What you should say out loud:**
- "I keep left/right pointers and move the one that helps satisfy the condition."
- "Invariant: everything left of `left` and right of `right` is already ruled out / processed."
- "This avoids nested loops → O(n)."

---

## 1) Overview
Two pointers uses two indices to traverse arrays/strings/linked lists efficiently. Often used on sorted arrays, palindrome checks, and partitioning problems.

## 2) Problem(s) it solves
- Pair sum in sorted array
- In-place deduplication on sorted array
- Palindrome checks
- Merge two sorted arrays
- Fast/slow pointers for cycle detection
- Container with most water

## 3) Core idea & invariants
- Opposite direction: left starts at 0, right at n-1.
- Same direction (fast/slow): fast explores, slow lags to build result.
- Linked list cycle: fast moves 2 steps, slow 1 step.

## 4) Typical formulations
- "Two sum in sorted array"
- "Remove duplicates in-place"
- "Detect cycle in linked list"
- "Max area container"

## 5) Complexity
- **Time:** O(n)
- **Space:** O(1) (excluding output)

## 6) Go notes
- Strings: be explicit whether you treat as bytes (ASCII) or runes (Unicode).
- In-place operations: return new length; don’t allocate unless required.

## 7) Pitfalls
1. Off-by-one loops.
2. Not moving pointer in some branches → infinite loop.
3. Assuming sorted input when it isn’t.

## 8) Practice katas
- Two Sum II (LC #167)
- Remove Duplicates (LC #26)
- Linked List Cycle (LC #141)
- Container With Most Water (LC #11)

## 9) Variations
- Three pointers (3Sum)
- Partitioning (Dutch national flag)
- Trapping rain water

## 10) Further reading
- https://leetcode.com/tag/two-pointers/
- https://go.dev/doc/effective_go

---

## Contract (your Go API)

```go
package two_pointers

func TwoSumSorted(nums []int, target int) []int // return 1-based indices, empty slice if not found
func RemoveDuplicates(nums []int) int           // in-place, return new length
func IsPalindrome(s string) bool                // define ASCII vs Unicode; tests assume ASCII
func ReverseArray(nums []int)
func MergeSortedArrays(a, b []int) []int

type ListNode struct {
	Val  int
	Next *ListNode
}
func HasCycleLinkedList(head *ListNode) bool
func ContainerWithMostWater(height []int) int
```