# Merge Intervals

## 0) Interview framing (Google/FAANG)
**What they test:** sorting + single pass merge, correct overlap logic, and edge cases.  
**What you should say out loud:**
- "Sort by start.”
- "Maintain a merged list; compare current interval with last merged.”
- "If overlapping, extend end; else append.”

---

## 1) Overview
Merge Intervals consolidates overlapping ranges. It appears in scheduling, calendar merging, log compaction, and range union problems.

## 2) Problems it solves
- Merge overlapping time blocks
- Insert interval into a merged set
- Compute minimum rooms for meetings

## 3) Core idea & invariants
- Sort intervals by `Start` (then `End`).
- Invariant: `merged` is always non-overlapping and sorted.
- Overlap condition (inclusive): `curr.Start <= last.End`.

## 4) Typical prompts
- "Merge intervals”
- "Insert interval”
- "Meeting rooms (min number of rooms)”

## 5) Complexity
- **Time:** O(n log n) due to sorting
- **Space:** O(n) for output

## 6) Go notes
- Use `sort.Slice`.
- Avoid in-place surprises unless documented.
- For MeetingRooms: min-heap of end times is the standard interview answer.

## 7) Pitfalls
1. Off-by-one overlap: inclusive endpoints.
2. Input not sorted.
3. Empty input / single interval.

## 8) Practice katas
- Merge Intervals (LC #56)
- Insert Interval (LC #57)
- Meeting Rooms II (LC #253)

## 9) Variations
- Exclusive endpoints
- Return total covered length
- Remove covered intervals / non-overlapping selection

## 10) Further reading
- https://leetcode.com/problems/merge-intervals/
- https://leetcode.com/problems/insert-interval/
- https://teddysmith.io/meeting-rooms-ii/

---

## Contract (your Go API)

```go
package merge_intervals

type Interval struct {
	Start int
	End   int
}

func MergeIntervals(intervals []Interval) []Interval
func InsertInterval(intervals []Interval, newInterval Interval) []Interval
func MeetingRooms(intervals []Interval) int
```