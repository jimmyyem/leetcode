package answers

//https://leetcode-cn.com/problems/next-permutation/
//31. 下一个排列
func NextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	//1. 找到反转点，就是找到要交换的小数。 123465 这里就是找到4
	//找到4后，然后找到后面比4大的交换位置 才能找到比当前大，但是幅度最小的
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	//如果i==-1则说明原始数字是最大的，比如 654321，会跳过下一步搜索123465后比4大的数字过程，而直接进入排序步骤3

	//2.从后往前找到比4大的第一个数，然后交换。123465 j这里最后停留在5这里
	//交换后是123564
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	//3. 把5后面的数字重新从小到大排序，即是最小的比原始数字大的符合条件数字
	//因为 123564 还是比 123546要大，而 123546 是符合条件的
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
