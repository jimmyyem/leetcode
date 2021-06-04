package answers

//https://leetcode-cn.com/problems/contiguous-array/
//525. 连续数组
func FindMaxLength(nums []int) int {
	maxLength := 0
	for i := 1; i < len(nums); {
		if nums[i] != nums[i-1] {
			maxLength += 2
			i += 2
			continue
		}
		i++
	}
	return maxLength
}
