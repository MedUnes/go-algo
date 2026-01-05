# Depth-First Search (DFS)

## 0) Interview framing (Google/FAANG)
**What they test:** recursion/stack discipline, visited invariants, cycle detection, and “explore then backtrack” reasoning.  
**What you should say out loud:**
- “DFS explores one branch fully before trying the next.”
- “Invariant: visited prevents infinite loops.”
- “For cycle detection in directed graphs, track recursion stack (colors).”

---

## 1) Overview
Depth-First Search traverses a graph by going as deep as possible before backtracking. DFS underpins cycle detection, topological sort (DFS postorder), SCC algorithms, and backtracking.

## 2) Problem(s) it solves
- Graph traversal / reachability
- Cycle detection (directed with recursion stack)
- Enumerating all paths (on DAG / with constraints)
- Backtracking-style search

## 3) Core idea & invariants
### Invariants
- **Visited invariant:** each node is processed once (for traversal).
- **Stack invariant (directed cycle):** a back-edge to a node “in stack” implies a cycle.
- **Path invariant (all-paths):** current path reflects recursion stack and is restored on return.

### Why it works
Following edges recursively explores the entire reachable subgraph; backtracking restores state enabling exploration of alternate branches.

## 4) Typical formulations (interview prompts)
- “Return DFS traversal order from node s.”
- “Detect if directed graph has a cycle.”
- “Return all paths from s to t (typically DAG).”

## 5) Complexity
- **Time:** O(V + E) for traversal/cycle detection
- **Space:** O(V) for visited + recursion/explicit stack
- All-paths can be exponential in number of paths.

## 6) Go implementation notes (idiomatic)
- Prefer iterative DFS if recursion depth might be large.
- If you implement both, ensure they match by enforcing deterministic neighbor order.
- For cycle detection: use 3-color (0=unvisited,1=visiting,2=done) or recursion-stack set.

## 7) Common pitfalls
1. Stack overflow on deep graphs (recursive DFS).
2. Mistaking undirected vs directed cycle detection logic.
3. For all-paths: forgetting to copy the current path before storing.

## 8) Practice katas
- Number of Islands (LC #200) (grid DFS)
- Course Schedule (LC #207) (cycle detection)
- All Paths From Source to Target (LC #797)
- Word Search (LC #79)

## 9) Variations you should recognize
- Preorder/postorder DFS
- DFS timestamps (discovery/finish)
- Tarjan/Kosaraju SCC
- DFS for topological sorting

## 10) Further reading
- https://walkccc.me/CLRS/Chap22/22.3/
- https://visualgo.net/en/dfsbfs

---

## Contract (your Go API)

```go
package dfs

type Graph struct {
	// your representation
}

func NewGraph() *Graph
func (g *Graph) AddEdge(u, v int)

func DFSRecursive(g *Graph, start int) []int
func DFSIterative(g *Graph, start int) []int

// HasCycle detects a cycle in a directed graph.
func HasCycle(g *Graph) bool

// FindAllPaths returns all simple paths from start to target.
// If you allow cycles, define how you prevent infinite paths (usually disallow revisiting nodes in current path).
func FindAllPaths(g *Graph, start, target int) [][]int
```