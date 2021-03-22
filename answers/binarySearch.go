package answers

//https://leetcode-cn.com/problems/binary-search/
//704.二分查找
func BinarySearch(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	var start, mid, end = 0, (n - 1) / 2, n - 1
	for start <= end {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
		mid = (start + end) / 2
	}

	return -1
}
