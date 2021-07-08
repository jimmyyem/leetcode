package answers

//https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
//154. 寻找旋转排序数组中的最小值
func FindMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) >> 1

		if nums[left] == nums[right] { //出现重复的元素直接略过
			left++
			continue
		}
		if nums[left] < nums[right] { //第一个出现转折点的地方
			return nums[left]
		}

		if nums[left] <= nums[mid] { //递增区间
			left = mid + 1
		} else { //包含旋转点
			right = mid
		}
	}

	return nums[left]
}
