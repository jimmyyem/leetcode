package answers

//https://leetcode-cn.com/problems/total-hamming-distance/
//477. 汉明距离总和
func TotalHammingDistance(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += val >> i & 1
		}
		ans += c * (n - c)
	}
	return
}
