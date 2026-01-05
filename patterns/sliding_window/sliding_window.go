package sliding_window

// MaxSumSubarray finds max sum of k consecutive elements.
// Return 0 if k <= 0 or k > len(nums).
func MaxSumSubarray(nums []int, k int) int

// LongestSubstringKDistinct returns the length of the longest substring with at most k distinct characters.
// Return 0 if k <= 0.
func LongestSubstringKDistinct(s string, k int) int

// MinWindowSubstring returns the minimum window in s which contains all characters of t (including multiplicity).
// Return "" if no such window exists.
func MinWindowSubstring(s string, t string) string
