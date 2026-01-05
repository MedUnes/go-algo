package sliding_window

import "testing"

// NOTE: No solutions here. These tests define the expected behavior.
// Implement the functions in sliding_window.go.

func TestMaxSumSubarray_Core(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"classic", []int{2, 1, 5, 1, 3, 2}, 3, 9}, // 5+1+3
		{"k_equals_1", []int{5, -1, 3}, 1, 5},
		{"all_negative", []int{-2, -1, -5}, 2, -3},
		{"k_equals_len", []int{1, 2, 3}, 3, 6},
		{"invalid_k_zero", []int{1, 2, 3}, 0, 0},
		{"invalid_k_too_big", []int{1, 2, 3}, 4, 0},
		{"empty_nums", []int{}, 1, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MaxSumSubarray(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("MaxSumSubarray(%v, %d): want %d, got %d", tt.nums, tt.k, tt.want, got)
			}
		})
	}
}

func TestLongestSubstringKDistinct_Core(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"eceba_k2", "eceba", 2, 3}, // "ece"
		{"aa_k1", "aa", 1, 2},       // "aa"
		{"abc_k2", "abc", 2, 2},     // "ab"/"bc"
		{"empty", "", 2, 0},
		{"k_zero", "abc", 0, 0},
		{"k_large", "abc", 10, 3},
		{"repeats", "aabacbebebe", 3, 7}, // common interview example: "cbebebe"
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := LongestSubstringKDistinct(tt.s, tt.k)
			if got != tt.want {
				t.Fatalf("LongestSubstringKDistinct(%q, %d): want %d, got %d", tt.s, tt.k, tt.want, got)
			}
		})
	}
}

func TestMinWindowSubstring_Core(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		tgt  string
		want string
	}{
		{"lc_example", "ADOBECODEBANC", "ABC", "BANC"},
		{"exact_match", "a", "a", "a"},
		{"no_solution", "a", "aa", ""},
		{"needs_multiplicity", "aaab", "aab", "aab"}, // multiplicity matters
		{"t_empty", "abc", "", ""},                   // define: empty target => empty window
		{"s_empty", "", "a", ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MinWindowSubstring(tt.s, tt.tgt)
			if got != tt.want {
				t.Fatalf("MinWindowSubstring(%q, %q): want %q, got %q", tt.s, tt.tgt, tt.want, got)
			}
		})
	}
}

// TODO (edge cases you should implement yourself):
// - Unicode/runes vs bytes: define the constraint and test accordingly.
// - Very large s (1e6) performance smoke test.
// - Multiple minimal windows: ensure you return the correct one per your spec (LC says unique in tests).
// - MinWindowSubstring: target has chars not in s at all.
// - LongestSubstringKDistinct: k==1 with alternating chars.
// - MaxSumSubarray: overflow behavior (use int64?) if constraints are huge.

// Bonus TODOs (variations):
// - SlidingWindowMaximum(nums, k) using a monotonic deque (LC #239).
// - LongestSubstringWithoutRepeating (LC #3).
// - SubarraysWithKDistinct (exactly k distinct).
