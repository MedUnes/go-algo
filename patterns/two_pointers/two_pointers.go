package two_pointers

func TwoSumSorted(nums []int, target int) []int // return 1-based indices, empty slice if not found
func RemoveDuplicates(nums []int) int           // in-place, return new length
func IsPalindrome(s string) bool                // define ASCII vs Unicode; tests assume ASCII
func ReverseArray(nums []int)
func MergeSortedArrays(a, b []int) []int

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycleLinkedList(head *ListNode) bool
func ContainerWithMostWater(height []int) int
