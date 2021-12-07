package answers

// https://leetcode-cn.com/problems/find-the-duplicate-number/
// 287. 寻找重复数  1，3，4，2，2
func FindDuplicate(nums []int) int {
	left, right := 0, len(nums)-1

	//o(logn)
	for left < right {
		mid, count := (left+right)/2, 0
		//o(n)
		for _, val := range nums {
			if val <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

/**

i=0 => 1
i=1 => 3
i=2 => 4
i=3 => 2
i=4 => 2

*/
