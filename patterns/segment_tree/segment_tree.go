package segment_tree

type SegmentTree struct {
	tree []int
	n    int
}

func NewSegmentTree(nums []int) *SegmentTree
func (st *SegmentTree) Update(index int, val int)
func (st *SegmentTree) RangeSum(left int, right int) int
func (st *SegmentTree) RangeMin(left int, right int) int
