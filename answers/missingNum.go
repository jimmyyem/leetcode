package answers

import (
	"sort"
)

//https://leetcode-cn.com/problems/missing-number/
//268. 丢失的数字
func MissingNum(nums []int) int {
	sort.Ints(nums)

	n := len(nums)

	//处理1 或者0  这样的情况
	if n == 1 {
		if nums[0] == 0 {
			return nums[0] + 1
		} else {
			return nums[0] - 1
		}
	}

	//处理这样的不连续情况，0,2,3
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] != 1 {
			return nums[i-1] + 1
		}
	}

	//处理 0,1,2  这样的连续情况
	if nums[0] == 0 {
		return nums[n-1] + 1
	} else {
		//处理 1，2，3 这样的情况
		return nums[0] - 1
	}
}
