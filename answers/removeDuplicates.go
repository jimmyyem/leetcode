package answers

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
//26. 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 0                        //慢指针，从0开始
	for i := 1; i < len(nums); i++ { //快指针，从1开始
		if nums[slow] != nums[i] {
			slow++
			nums[slow] = nums[i]
		}
	}
	nums = nums[:slow+1]

	return slow + 1
}

/**

before [0 0 1 1 1 2 2 3 3 4]
before [0 0 1 1 1 2 2 3 3 4]
after [0 1 1 1 1 2 2 3 3 4]
before [0 1 1 1 1 2 2 3 3 4]
before [0 1 1 1 1 2 2 3 3 4]
before [0 1 1 1 1 2 2 3 3 4]
after [0 1 2 1 1 2 2 3 3 4]
before [0 1 2 1 1 2 2 3 3 4]
before [0 1 2 1 1 2 2 3 3 4]
after [0 1 2 3 1 2 2 3 3 4]
before [0 1 2 3 1 2 2 3 3 4]
before [0 1 2 3 1 2 2 3 3 4]
after [0 1 2 3 4 2 2 3 3 4]
5


*/
