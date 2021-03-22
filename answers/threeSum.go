package answers

import (
	"sort"
)

//https://leetcode-cn.com/problems/3sum/
//15. 三数之和
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}

	//排序
	sort.Ints(nums)

	//固定第一个数字
	for i := 0; i < len(nums); i++ {
		//第一个书就大于0则结果一定不等于0
		if nums[i] > 0 {
			break
		}

		//只要是不是第一个数，任何与他前面数相同（即1st 2nd数字相同则3rd一定也相同）
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			switch {
			case sum == 0:
				res = append(res, []int{nums[i], nums[start], nums[end]})
				//这里目的是固定i和end的情况把其他相同的start都给略过(-1,-1,-1,0,1,2)
				for start < end && nums[start+1] == nums[start] {
					start++
				}
				//这里目的是固定i和start的情况把其他相同的end都给略过(-1,0,1,1,1)
				for start < end && nums[end-1] == nums[end] {
					end--
				}

				//这里是最后再++和--，因为上面2步要把相同值的start和end给略过
				start++
				end--
			case sum > 0:
				end--
			default:
				start++
			}
		}
	}
	return res
}
