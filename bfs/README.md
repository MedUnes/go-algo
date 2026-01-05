# Breadth-First Search (BFS)

## 0) Interview framing (Google/FAANG)
**What they test:** graph modeling, shortest path in unweighted graphs, queue discipline, and visited invariants.  
**What you should say out loud:**
- “BFS explores level-by-level using a queue.”
- “Invariant: when a node is dequeued, we have the shortest distance to it (unweighted graph).”
- “We mark visited when enqueueing to avoid duplicates and cycles.”

---

## 1) Overview
Breadth-First Search traverses a graph in increasing distance from a start node. It’s the default tool for shortest paths in **unweighted** graphs and “k hops away” problems.

## 2) Problem(s) it solves
- Traverse all nodes reachable from a source
- Compute shortest path (fewest edges) in unweighted graphs
- Build level/layer partitions (distance buckets)

## 3) Core idea & invariants
### Invariants
- **Visited invariant:** each node enters the queue at most once.
- **Distance invariant:** first time we reach a node is via the shortest number of edges.
- **Level invariant (if using levels):** nodes in the same level share equal distance from start.

### Why it works (high level)
Queue order ensures we process all nodes at distance `d` before any node at distance `d+1`.

## 4) Typical formulations (interview prompts)
- “Return BFS traversal order starting from node s.”
- “Return nodes grouped by distance (levels).”
- “Return shortest path between s and t in an unweighted graph.”

## 5) Complexity
- **Time:** O(V + E)
- **Space:** O(V) (queue + visited + parent/dist maps)

## 6) Go implementation notes (idiomatic)
- Use adjacency list (map or slice-of-slices) and a slice-based queue (`[]int` with head index) to avoid `container/list` overhead.
- Mark visited on enqueue, not dequeue.
- For shortest path reconstruction, maintain `parent` map/slice and rebuild by walking parents backward.

## 7) Common pitfalls (what interviewers look for)
1. Marking visited too late → duplicates and possible blow-ups on dense graphs.
2. Assuming graph is connected.
3. Not defining behavior for start node not present / empty graph.

## 8) Practice katas (remember the pattern)
- Binary Tree Level Order Traversal (LC #102)
- Shortest Path in Binary Matrix (LC #1091)
- Rotting Oranges (LC #994)
- Word Ladder (LC #127)

## 9) Variations you should recognize
- Multi-source BFS
- Bidirectional BFS
- BFS on grid / implicit graph
- BFS that stops early when target found

## 10) Further reading
- https://walkccc.me/CLRS/Chap22/22.2/
- https://visualgo.net/en/dfsbfs

---

## Contract (your Go API)
Implement these functions/types in this package (exact names used by tests):

```go
package bfs

type Graph struct {
	// your representation
}

func NewGraph() *Graph
func (g *Graph) AddEdge(u, v int)
func (g *Graph) AddUndirectedEdge(u, v int)

func BFS(g *Graph, start int) []int
func BFSWithLevels(g *Graph, start int) [][]int

// ShortestPath returns (path, distance). If no path: (nil, -1).
func ShortestPath(g *Graph, start, target int) ([]int, int)
```