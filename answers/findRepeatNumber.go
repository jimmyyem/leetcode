package answers

//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func FindRepeatNumber(nums []int) int {
	//var list = make([]int, len(nums))
	//
	//for _, val := range nums {
	//	list[val]++
	//}
	//
	//for key, val := range list {
	//	if val > 1 {
	//		return key
	//	}
	//}

	var i int
	for i < len(nums) {
		//索引i的值是i，继续下一位
		if nums[i] == i {
			i++
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}

		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}

	return -1
}

/**
	2	3	1	0	2	5	3
------------------------------------
	1  	3 	2	0	2	5	3
	3	1	2	0	2	5	3
	0	1	2	3	2	5	3


 */