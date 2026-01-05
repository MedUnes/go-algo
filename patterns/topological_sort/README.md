# Topological Sort

## 0) Interview framing (Google/FAANG)
**What they test:** DAG ordering, cycle detection, and correctness proof via in-degree / DFS postorder.  
**What you should say out loud:**
- "Topological ordering exists iff the graph is acyclic.”
- "Kahn’s algorithm repeatedly removes nodes with in-degree 0.”
- "I verify cycles by checking whether I processed all nodes.”

---

## 1) Overview
Topological Sort produces a linear ordering of vertices in a **DAG** such that for every directed edge `u -> v`, `u` appears before `v`.

## 2) Problems it solves
- Build systems / compilation order
- Dependency resolution
- Course schedule feasibility and ordering

## 3) Core idea & invariants
### Kahn (BFS / in-degree)
- Invariant: queue contains only nodes with in-degree 0 (no unmet prerequisites).
- Each time you pop `u`, you "remove” its outgoing edges by decrementing neighbors’ in-degrees.

### DFS-based
- Invariant: pushing node to output on postorder (after exploring children) yields reverse topological order.
- Cycle detection uses recursion stack / color marking.

## 4) Typical interview prompts
- "Return any topological order of nodes 0..n-1.”
- "Detect if cycle exists.”
- "Course Schedule / Course Schedule II.”

## 5) Complexity
- **Time:** O(V + E)
- **Space:** O(V + E) for adjacency + stacks/queues

## 6) Go notes (idiomatic)
- Use adjacency list: `[][]int` or `map[int][]int`.
- Keep in-degree in `[]int`.
- Don’t over-abstract graphs in interviews. Keep structs small.

## 7) Common pitfalls
1. Assuming a unique order (it’s not).
2. Forgetting isolated nodes (no edges).
3. Wrong direction of edges for prerequisites.

## 8) Practice katas
- Course Schedule (LC #207)
- Course Schedule II (LC #210)
- Alien Dictionary (LC #269)

## 9) Variations
- Detect cycle and return one cycle.
- Multiple valid orders: lexicographically smallest order (use heap).
- Build layers/levels (like semesters).

## 10) Further reading
- https://walkccc.me/CLRS/Chap22/22.4/
- https://www.geeksforgeeks.org/dsa/topological-sorting-indegree-based-solution/

---

## Contract (your Go API)

```go
package topological_sort

type DirectedGraph struct {
	// your representation
}

func NewDirectedGraph(n int) *DirectedGraph
func (g *DirectedGraph) AddEdge(u, v int)

func (g *DirectedGraph) NumVertices() int
func (g *DirectedGraph) Edges() [][2]int // optional helper for tests; you can expose differently

func TopologicalSortKahn(g *DirectedGraph) []int
func TopologicalSortDFS(g *DirectedGraph) []int
func CanFinishCourses(numCourses int, prerequisites [][]int) bool
```