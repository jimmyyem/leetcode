package answers

// https://leetcode-cn.com/problems/counting-bits/
// 338. 比特位计数
func CountBits(n int) []int {
	res := make([]int, 0, n+1)
	res = append(res, 0)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			res = append(res, res[i/2])
		} else {
			res = append(res, res[i-1]+1)
		}
	}

	return res
}
