package answers

//获得最长子序列的长度
//动态规划
func GetAllPath(nums []int) int {
	return getPath(nums, 0)
}

// 1, 5, 2, 4, 3
func getPath(nums []int, start int) int {
	maxLength := 1
	for i := start + 1; i < len(nums); i++ {
		if nums[start] < nums[i] {
			maxLength = max(maxLength, getPath(nums, i)+1)
		}
	}

	return maxLength
}
