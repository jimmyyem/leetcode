package answers

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)/2, 0, 0
	for low < high {
		nums1Mid = low + (high-low)>>1
		nums2Mid = k - nums1Mid

		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {

		}
	}

	return float64(1)
}

func min(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return y
	}
	return y
}
