
# Sliding Window

## 0) Interview framing (Google/FAANG)
**What they test:** Maintaining a rolling state (sum/frequency) while moving two pointers.  
**What you should say out loud:**
- "I maintain a window `[l..r]` and a state (sum / counts / distinct).”
- "I expand `r` until valid, then shrink `l` to minimal while staying valid.”
- "Invariant: state always matches the current window.”

---

## 1) Overview
Sliding Window turns many "all subarrays/substrings” problems from O(n²) brute force into O(n) by maintaining a dynamic window and updating a small state incrementally.

## 2) Problems it solves
- Fixed-size window aggregations (max sum / average)
- Longest/shortest substring satisfying constraints
- Minimum covering window (e.g., contains all chars of `t`)

## 3) Core idea & invariants
### Invariants
- Window is `[left, right]` (or `[left, right)`); define it once and stick to it.
- State (sum / frequency map / distinct count) always reflects exactly the window content.
- For "min window”: `formed == required` means the window is valid.

### Why it works
You never move pointers backward, so each character/element is processed O(1) times on average → total O(n).

## 4) Typical formulations
- "Max sum of any subarray of size k”
- "Longest substring with at most k distinct chars”
- "Minimum substring containing all chars of t” (Minimum Window Substring)

## 5) Complexity
- **Time:** O(n) typical (each pointer moves at most n)
- **Space:** O(k) / O(Σ) for frequency tracking (charset size)

## 6) Go notes (idiomatic)
- For ASCII problems, you can use `[256]int` counts; otherwise use `map[rune]int` and state your assumption.
- Don’t slice strings per step (allocations). Prefer indices and counts.
- Be explicit about:
  - returning `0` / `-1` / `""` on invalid input (tests cover this).

## 7) Common pitfalls
1. Off-by-one with window boundaries.
2. Forgetting to decrement counts when moving `left`.
3. Not defining "k distinct” as **at most** vs **exactly**.

## 8) Practice katas
- Longest Substring Without Repeating Characters (LC #3)
- Minimum Window Substring (LC #76)
- Sliding Window Maximum (LC #239)

## 9) Variations
- Fixed vs variable window
- "At most k” vs "exactly k” (often derived)
- Monotonic queue (for max/min in window)

## 10) Further reading
- https://leetcode.com/problems/minimum-window-substring/
- https://leetcode.com/problems/sliding-window-maximum/
- https://labuladong.online/algo/en/essential-technique/sliding-window-framework/

---

## Contract (your Go API)

```go
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
```