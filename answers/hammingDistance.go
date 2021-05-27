package answers

//https://leetcode-cn.com/problems/hamming-distance/
//461. 汉明距离
func HammingDistance(x int, y int) int {
	//return bits.OnesCount(uint(x) ^ uint(y))

	res := 0
	n := x ^ y
	for n != 0 {
		res++
		n = n & (n-1)
	}
	return res
}
