package merge_intervals

type Interval struct {
	Start int
	End   int
}

func MergeIntervals(intervals []Interval) []Interval
func InsertInterval(intervals []Interval, newInterval Interval) []Interval
func MeetingRooms(intervals []Interval) int
