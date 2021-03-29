package answers

//https://leetcode-cn.com/problems/remove-element/
//27. 删除元素
func RemoveElement(nums []int, val int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[slow], nums[i] = nums[i], nums[slow]
			slow++
		}
	}

	return slow
}

/**

slow 	i   data
0	 	0	3 2 2 3     i++
0		1   2 3 2 3     i++ slow++ = 1
1		2   2 2 3 3		i++ slow++  = 2
2       3   2 2 3 3		i++
*/
