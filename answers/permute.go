package answers

// 最终结果
var result [][]int

//https://leetcode-cn.com/problems/permutations
//46.全排列

func Permute(nums []int) [][]int {
	var pathNums = make([]int, 0) //这里后面要根据len判断是否到达最末级，所以最好初始化为0个元素
	var used = make([]bool, len(nums))
	result = [][]int{} //reset全局变量（leetcode多次执行全局变量不会消失）
	backtrack(nums, pathNums, used)
	return result
}

// 回溯核心
// nums: 原始列表
// pathNums: 路径上的数字
// used: 是否访问过
func backtrack(nums, pathNums []int, used []bool) {
	// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, pathNums)
		// 把本次结果追加到最终结果上
		result = append(result, tmp)
		return
	}

	// 开始遍历原始数组的每个数字
	for i := 0; i < len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			// 撤销刚才的选择，也就是恢复操作
			pathNums = pathNums[:len(pathNums)-1]
			// 标记成未使用
			used[i] = false
		}
	}
}
