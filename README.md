# go-algo

[![Go Report Card](https://goreportcard.com/badge/github.com/medunes/go-algo)](https://goreportcard.com/report/github.com/medunes/go-algo)
[![Go Reference](https://pkg.go.dev/badge/github.com/medunes/go-algo.svg)](https://pkg.go.dev/github.com/medunes/go-algo)
[![Tests](https://github.com/medunes/go-algo/actions/workflows/tests.yml/badge.svg)](https://github.com/medunes/go-algo/actions/workflows/tests.yml)

![image](./algorithms.png)

A focused collection of algorithm patterns implemented in **idiomatic Go**, designed for **FAANG-style interviews** and real-world intuition.  
Each folder is self-contained and includes:
- a concise README (what it solves, invariants, complexity, common pitfalls, katas)
- `_test.go` files with **core tests** plus **TODO edge-case tests** you implement yourself

## Contents

- [Breadth-First Search (BFS)](./bfs)
- [Depth-First Search (DFS)](./dfs)
- [Two Heaps (Median / Priority pattern)](./two_heaps)
- [Two Pointers](./two_pointers)
- [Sliding Window](./sliding_window)
- [Topological Sort](./topological_sort)
- [Merge Intervals](./merge_intervals)
- [Backtracking](./back_tracking)
- [Trie (Prefix Tree)](./trie)
- [Flood Fill (Grid BFS/DFS)](./flood_fill)
- [Segment Tree (Range queries)](./segment_tree)

## How to use this repo

1. Pick a folder.
2. Read the README (focus on invariants and typical interview formulations).
3. Implement the missing functions (no solution code is provided by design).
4. Run tests, then implement the TODO edge-case tests for mastery.

## Running tests

Run everything:
```bash
make test
go run gotest.tools/gotestsum@latest --format=testdox
github.com/medunes/go-algo/bfs:
 ✓ BFS core traversal (0.00s)
 ✓...
DONE 36 tests in 0.081s
```
