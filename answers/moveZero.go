package answers

//https://leetcode-cn.com/problems/move-zeroes/
//283. 移动零
func MoveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	slow, fast, size := 0, 0, len(nums)
	for fast < size {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
