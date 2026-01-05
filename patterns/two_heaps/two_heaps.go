package two_heaps

type MedianFinder struct {
	// your fields
}

func NewMedianFinder() *MedianFinder
func (mf *MedianFinder) AddNum(num int)
func (mf *MedianFinder) FindMedian() float64

type MaxHeap []int
type MinHeap []int
