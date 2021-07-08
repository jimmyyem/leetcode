package answers

//https://leetcode-cn.com/problems/total-hamming-distance/
//477. 汉明距离总和
func TotalHammingDistance(nums []int) int {
	var n, ans = len(nums), 0
	for i := 0; i < 30; i++ {
		c := 0 //记录每次循环当前位（第i位）是1的个数
		for _, val := range nums {
			//利用与运算，如果结果是0，说明当前位相同，加上0也白加；如果结果是1则直接计算；少了一次if判断
			//c += (val >> i) & 1

			//利用计算余数更简单易懂
			if (val>>i)%2 == 1 {
				c++
			}
		}
		ans += c * (n - c)
	}

	return ans
}

/**
https://www.bilibili.com/video/BV1SW411r78m

从1到31位，挨个遍历每一位，每一位的值要么0，要么1
1的个位数cnt1，0的个数为cnt0， cnt0*cnt1 即为该位上汉明距离总和
...
31

这样31个数相加即为汉明距离总和
*/
