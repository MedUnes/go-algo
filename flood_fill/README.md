# Flood Fill

## 0) Interview framing (Google/FAANG)
**What they test:** grid traversal, visited marking, and BFS/DFS choice.  
**What you should say out loud:**
- “This is graph traversal on a grid.”
- “Neighbors are 4-directional (unless stated otherwise).”
- “Use BFS/DFS; mark visited to avoid loops.”

---

## 1) Overview
Flood Fill recolors/marks a connected region in a grid. The same pattern solves “count islands” and “max island area”.

## 2) Problems it solves
- Flood Fill (recolor region)
- Number of Islands
- Max Area of Island

## 3) Core invariants
- Only cells connected via valid directions are included.
- Visited cells are not processed twice.

## 4) Complexity
- **Time:** O(R*C)
- **Space:** O(R*C) worst-case for queue/stack

## 5) Go notes
- Avoid recursion depth issues on large grids (iterative BFS is safer).
- Be explicit about bounds checks.
- Decide whether you mutate in place or use visited.

## 6) Pitfalls
1. Not handling `newColor == oldColor` in FloodFill (should no-op).
2. Mixing byte grids and int grids incorrectly.
3. Forgetting to mark visited early → repeated work.

## 7) Practice katas
- Flood Fill (LC #733)
- Number of Islands (LC #200)
- Max Area of Island (LC #695)

## 8) Further reading
- https://www.geeksforgeeks.org/dsa/flood-fill-algorithm/
- https://leetcode.com/problems/number-of-islands/
- https://leetcode.com/problems/flood-fill/

---

## Contract (your Go API)

```go
package flood_fill

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int
func NumIslands(grid [][]byte) int
func MaxAreaOfIsland(grid [][]int) int
```