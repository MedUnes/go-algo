package two_pointers

import "testing"

func TestTwoSumSorted_Core(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"found", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"not_found", []int{1, 2, 3, 4}, 10, []int{}},
		{"first_last", []int{1, 2, 3, 4, 5}, 6, []int{1, 5}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := TwoSumSorted(tt.nums, tt.target)
			assertIntSliceEqual(t, tt.want, got)
		})
	}
}

func TestRemoveDuplicates_Core(t *testing.T) {
	t.Parallel()

	nums := []int{1, 1, 2, 2, 3, 4, 4}
	n := RemoveDuplicates(nums)

	wantPrefix := []int{1, 2, 3, 4}
	if n != len(wantPrefix) {
		t.Fatalf("RemoveDuplicates length: want %d, got %d (nums=%v)", len(wantPrefix), n, nums)
	}
	for i := 0; i < n; i++ {
		if nums[i] != wantPrefix[i] {
			t.Fatalf("RemoveDuplicates prefix[%d]: want %d, got %d (nums=%v)", i, wantPrefix[i], nums[i], nums)
		}
	}

	nums2 := []int{1, 2, 3}
	n2 := RemoveDuplicates(nums2)
	if n2 != 3 {
		t.Fatalf("RemoveDuplicates no-dups: want 3, got %d", n2)
	}

	nums3 := []int{1, 1, 1, 1}
	n3 := RemoveDuplicates(nums3)
	if n3 != 1 || nums3[0] != 1 {
		t.Fatalf("RemoveDuplicates all-same: want len=1 prefix=[1], got len=%d nums=%v", n3, nums3)
	}
}

func TestIsPalindrome_Core(t *testing.T) {
	t.Parallel()

	ok := []string{"racecar", "a", "aa", ""}
	for _, s := range ok {
		s := s
		t.Run("true_"+s, func(t *testing.T) {
			t.Parallel()
			if !IsPalindrome(s) {
				t.Fatalf("IsPalindrome(%q): want true, got false", s)
			}
		})
	}

	bad := []string{"hello", "ab", "abcd"}
	for _, s := range bad {
		s := s
		t.Run("false_"+s, func(t *testing.T) {
			t.Parallel()
			if IsPalindrome(s) {
				t.Fatalf("IsPalindrome(%q): want false, got true", s)
			}
		})
	}
}

func TestReverseArray_Core(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3, 4, 5}
	ReverseArray(a)
	assertIntSliceEqual(t, []int{5, 4, 3, 2, 1}, a)

	b := []int{1, 2, 3, 4}
	ReverseArray(b)
	assertIntSliceEqual(t, []int{4, 3, 2, 1}, b)
}

func TestMergeSortedArrays_Core(t *testing.T) {
	t.Parallel()

	got := MergeSortedArrays([]int{1, 3, 5}, []int{2, 4, 6})
	assertIntSliceEqual(t, []int{1, 2, 3, 4, 5, 6}, got)

	got2 := MergeSortedArrays([]int{1, 2, 3}, []int{})
	assertIntSliceEqual(t, []int{1, 2, 3}, got2)
}

func TestHasCycleLinkedList_Core(t *testing.T) {
	t.Parallel()

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	if HasCycleLinkedList(head) {
		t.Fatalf("HasCycleLinkedList: expected false, got true")
	}

	// create cycle
	h2 := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	third := &ListNode{Val: 3}
	h2.Next = second
	second.Next = third
	third.Next = second
	if !HasCycleLinkedList(h2) {
		t.Fatalf("HasCycleLinkedList: expected true, got false")
	}
}

func TestContainerWithMostWater_Core(t *testing.T) {
	t.Parallel()

	if got := ContainerWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}); got != 49 {
		t.Fatalf("ContainerWithMostWater: want 49, got %d", got)
	}
	if got := ContainerWithMostWater([]int{1, 2}); got != 1 {
		t.Fatalf("ContainerWithMostWater(2 elems): want 1, got %d", got)
	}
}

// --- helpers ---

func assertIntSliceEqual(t *testing.T, want, got []int) {
	t.Helper()
	if len(want) != len(got) {
		t.Fatalf("slice length mismatch: want=%d got=%d\nwant=%v\ngot=%v", len(want), len(got), want, got)
	}
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("slice[%d] mismatch: want=%d got=%d\nwant=%v\ngot=%v", i, want[i], got[i], want, got)
		}
	}
}

// TODO (edge cases you should implement yourself):
// - TwoSumSorted with negative numbers.
// - Empty / single-element arrays.
// - RemoveDuplicates empty array.
// - IsPalindrome with spaces/punctuation and case folding (define spec).
// - ReverseArray empty/single.
// - MergeSortedArrays both empty / very different sizes.
// - HasCycleLinkedList self-loop.
// - ContainerWithMostWater all same heights.
//
// Bonus TODOs:
// - ThreeSum.
// - Trapping rain water.
// - Dutch national flag.
