# Backtracking

## 0) Interview framing (Google/FAANG)
**What they test:** recursion structure, pruning, correctness, and avoiding global bugs.  
**What you should say out loud:**
- "I build a partial solution, and when it violates constraints, I stop exploring.”
- "I revert state before returning (undo step).”
- "Pruning reduces branching.”

---

## 1) Overview
Backtracking is DFS over the space of possibilities, with explicit "choose → explore → unchoose” steps.

## 2) Problems it solves
- Permutations / combinations / subsets
- Constraint satisfaction (N-Queens, Sudoku)
- Word search / path constraints

## 3) Core invariants
- The partial solution always satisfies constraints.
- State is fully restored after returning from recursion.

## 4) Typical prompts
- "Generate all permutations”
- "Choose k out of n”
- "Solve N-Queens”
- "Solve Sudoku”

## 5) Complexity
- Usually exponential (often `O(n!)` or `O(C(n,k))`)
- Space: recursion depth + output size

## 6) Go notes
- Be careful with slice aliasing: copy when storing solutions.
- Prefer local slices and append/pop pattern.
- Use boolean arrays for used/cols/diagonals in N-Queens for O(1) checks.

## 7) Pitfalls
1. Forgetting to undo state (classic bug).
2. Storing pointers/slices without copying.
3. Weak pruning → TLE.

## 8) Practice katas
- Permutations (LC #46)
- Combinations (LC #77)
- Sudoku Solver (LC #37)
- N-Queens (LC #51)

## 9) Further reading
- https://www.geeksforgeeks.org/dsa/n-queen-problem-backtracking-3/
- https://algo.monster/liteproblems/51

---

## Contract (your Go API)

```go
package backtracking

func Permutations(nums []int) [][]int
func Combinations(n int, k int) [][]int
func SolveSudoku(board [][]byte) bool
func NQueens(n int) [][]string
```