# Trie (Prefix Tree)

## 0) Interview framing (Google/FAANG)
**What they test:** prefix-based search, space/time tradeoffs, and clean API.  
**What you should say out loud:**
- "Each node represents a prefix; traversing chars is O(L).”
- "Insert/Search/StartsWith are O(L).”
- "Memory tradeoff: many nodes; can compress via radix tree.”

---

## 1) Overview
A Trie stores strings by shared prefixes, enabling fast prefix queries (autocomplete/spell-check) in time proportional to the word length.

## 2) Problems it solves
- Insert/search dictionary words
- Prefix queries
- Autocomplete results for a prefix

## 3) Core invariants
- Every node corresponds to a prefix.
- `isEnd` marks full word termination.
- Children contain transitions by character.

## 4) Complexity
- **Insert/Search/StartsWith:** O(L) where L is key length
- Space: depends on branching; can be large.

## 5) Go notes
- Decide byte vs rune:
    - ASCII only: `map[byte]*node` or `[26]*node`
    - Unicode: `map[rune]*node`
- Avoid recursion for `SearchWords` if you want iterative DFS.

## 6) Pitfalls
1. Forgetting multiplicity of paths in autocomplete traversal.
2. Confusing prefix existence vs word existence.
3. Returning unstable ordering (tests should not assume order).

## 7) Practice katas
- Implement Trie (LC #208)
- Design Add and Search Words (LC #211)
- Word Search II (LC #212)

## 8) Further reading
- https://en.wikipedia.org/wiki/Trie
- https://leetcode.com/problems/implement-trie-prefix-tree/

---

## Contract (your Go API)

```go
package trie

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie

func (t *Trie) Insert(word string)
func (t *Trie) Search(word string) bool
func (t *Trie) StartsWith(prefix string) bool
func (t *Trie) SearchWords(prefix string) []string // Autocomplete
```