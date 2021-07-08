package answers

import "sort"

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
//4. 寻找两个正序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merge := make([]int, 0, len(nums1)+len(nums2))
	merge = append(merge, nums1...)
	merge = append(merge, nums2...)
	sort.Ints(merge)
	if len(merge)%2 == 1 {
		return float64(merge[len(merge)/2])
	}
	return (float64(merge[len(merge)/2]) + float64(merge[len(merge)/2-1])) / 2

	/**
	1. nums1短，中位数在nums2里
	2. nums1长，中位数在nums1里
	3.
	*/


}
