package answers

//https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
//34. 在排序数组中查找元素的第一个和最后一个位置
func SearchRange(nums []int, target int) []int {
	first := lowerBound(nums, target)
	second := upperBound(nums, target)

	if first > second {
		return []int{-1, -1}
	}
	return []int{first, second}
}

//获取第一个出现的位置
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

//获取最后一个出现的位置
func upperBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
