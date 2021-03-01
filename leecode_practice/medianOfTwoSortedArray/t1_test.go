package medianOfTwoSortedArray

import "testing"

func TestA(t *testing.T)  {
	nums1 := []int{3}
	nums2 := []int{-2,-1}
	t.Log(findMedianSortedArrays(nums1,nums2))
}
