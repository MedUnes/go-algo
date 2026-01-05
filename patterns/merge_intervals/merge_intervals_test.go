package merge_intervals

import (
	"testing"
)

func TestMergeIntervals_Core(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []Interval
		want []Interval
	}{
		{
			"lc_example",
			[]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[]Interval{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"already_merged",
			[]Interval{{1, 2}, {3, 4}},
			[]Interval{{1, 2}, {3, 4}},
		},
		{
			"touching_edges_merge_inclusive",
			[]Interval{{1, 2}, {2, 3}},
			[]Interval{{1, 3}},
		},
		{
			"single",
			[]Interval{{5, 7}},
			[]Interval{{5, 7}},
		},
		{
			"empty",
			[]Interval{},
			[]Interval{},
		},
		{
			"unsorted_input",
			[]Interval{{8, 10}, {1, 3}, {2, 6}},
			[]Interval{{1, 6}, {8, 10}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MergeIntervals(tt.in)
			assertIntervalsEqual(t, tt.want, got)
		})
	}
}

func TestInsertInterval_Core(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []Interval
		add  Interval
		want []Interval
	}{
		{
			"insert_middle_merge",
			[]Interval{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			Interval{4, 8},
			[]Interval{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			"insert_no_overlap",
			[]Interval{{1, 2}, {5, 6}},
			Interval{3, 4},
			[]Interval{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			"insert_front",
			[]Interval{{5, 7}},
			Interval{1, 2},
			[]Interval{{1, 2}, {5, 7}},
		},
		{
			"insert_back",
			[]Interval{{1, 2}},
			Interval{5, 7},
			[]Interval{{1, 2}, {5, 7}},
		},
		{
			"empty_in",
			[]Interval{},
			Interval{2, 3},
			[]Interval{{2, 3}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := InsertInterval(tt.in, tt.add)
			assertIntervalsEqual(t, tt.want, got)
		})
	}
}

func TestMeetingRooms_Core(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []Interval
		want int
	}{
		{
			"lc_style",
			[]Interval{{0, 30}, {5, 10}, {15, 20}},
			2,
		},
		{
			"non_overlapping",
			[]Interval{{7, 10}, {2, 4}},
			1,
		},
		{
			"all_overlap",
			[]Interval{{1, 5}, {2, 6}, {3, 7}},
			3,
		},
		{
			"empty",
			[]Interval{},
			0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MeetingRooms(tt.in)
			if got != tt.want {
				t.Fatalf("MeetingRooms(%v): want %d, got %d", tt.in, tt.want, got)
			}
		})
	}
}

// --- helpers ---

func assertIntervalsEqual(t *testing.T, want, got []Interval) {
	t.Helper()

	if len(want) != len(got) {
		t.Fatalf("intervals length mismatch: want %d, got %d\nwant=%v\ngot=%v", len(want), len(got), want, got)
	}
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("interval[%d] mismatch: want %v, got %v\nwant=%v\ngot=%v", i, want[i], got[i], want, got)
		}
	}
}

// TODO (edge cases you should implement yourself):
// - Invalid intervals where Start > End: decide normalize vs error.
// - Very large n performance test.
// - Duplicates and contained intervals: [1,10] with [2,3].
// - MeetingRooms: intervals with same start time.
// - Boundary rules: if [1,2] and [2,3] share a room or not (define; tests assume YES reuse at start==end).
