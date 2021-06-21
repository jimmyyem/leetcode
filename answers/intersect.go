package answers

import (
	"sort"
)

//https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/solution/
//350. 两个数组的交集 II
func Intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	index1, index2 := 0, 0

	sort.Ints(nums1)
	sort.Ints(nums2)

	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] == nums2[index2] {
			result = append(result, nums1[index1])
			index1++
			index2++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			index1++
		}
	}

	return result
}

func Intersect2(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	cache := make(map[int]int)

	for _, val := range nums1 {
		cache[val]++
	}

	for _, val := range nums2 {
		if n, ok := cache[val]; ok && n > 0 {
			result = append(result, val)
			cache[val]--
		}
	}

	return result
}
