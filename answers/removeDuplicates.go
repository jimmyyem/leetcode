package answers

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[idx] != nums[i] {
			idx++
			nums[idx] = nums[i]
		}
	}
	//fmt.Println(idx, nums, nums[:idx+1])
	nums = nums[:idx+1]

	return idx + 1
}
