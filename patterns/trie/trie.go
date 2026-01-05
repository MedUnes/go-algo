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
