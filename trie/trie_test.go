package trie

import "testing"

func TestTrie_InsertAndSearch_Core(t *testing.T) {
	t.Parallel()

	tr := NewTrie()
	words := []string{"apple", "app", "apply", "bat", "bath"}

	for _, w := range words {
		tr.Insert(w)
	}

	tests := []struct {
		word string
		want bool
	}{
		{"apple", true},
		{"app", true},
		{"appl", false},
		{"bat", true},
		{"bath", true},
		{"batman", false},
		{"", false}, // define policy: empty word not stored unless explicitly inserted
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.word, func(t *testing.T) {
			t.Parallel()
			got := tr.Search(tt.word)
			if got != tt.want {
				t.Fatalf("Search(%q): want %v, got %v", tt.word, tt.want, got)
			}
		})
	}
}

func TestTrie_StartsWith_Core(t *testing.T) {
	t.Parallel()

	tr := NewTrie()
	tr.Insert("apple")
	tr.Insert("app")
	tr.Insert("banana")

	tests := []struct {
		prefix string
		want   bool
	}{
		{"app", true},
		{"ap", true},
		{"ban", true},
		{"bana", true},
		{"band", false},
		{"", true}, // define policy: empty prefix always true
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.prefix, func(t *testing.T) {
			t.Parallel()
			got := tr.StartsWith(tt.prefix)
			if got != tt.want {
				t.Fatalf("StartsWith(%q): want %v, got %v", tt.prefix, tt.want, got)
			}
		})
	}
}

func TestTrie_Autocomplete_Core(t *testing.T) {
	t.Parallel()

	tr := NewTrie()
	for _, w := range []string{"app", "apple", "apply", "apt", "bat"} {
		tr.Insert(w)
	}

	got := tr.SearchWords("app")

	// Order is NOT specified; treat as set.
	want := []string{"app", "apple", "apply"}
	assertStringSetEqual(t, want, got)
}

// --- helpers ---

func assertStringSetEqual(t *testing.T, want, got []string) {
	t.Helper()

	wm := map[string]int{}
	for _, w := range want {
		wm[w]++
	}
	gm := map[string]int{}
	for _, g := range got {
		gm[g]++
	}

	if len(wm) != len(gm) {
		t.Fatalf("set size mismatch: want=%v got=%v", want, got)
	}
	for k, v := range wm {
		if gm[k] != v {
			t.Fatalf("missing or wrong count for %q: want=%d got=%d (got=%v)", k, v, gm[k], got)
		}
	}
}

// TODO (edge cases you should implement yourself):
// - Insert same word twice: idempotent behavior.
// - Unicode words (runes): decide and test.
// - SearchWords for prefix that is itself a word and also has longer words.
// - Very deep single-branch trie (performance, stack usage).
// - Memory: large dictionary behavior (benchmark).
