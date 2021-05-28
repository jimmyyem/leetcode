package answers

//https://leetcode-cn.com/problems/hamming-distance/
//461. 汉明距离
func HammingDistance(x int, y int) int {
	res := 0
	for i := 0; i < 31; i++ {
		bx := x % 2
		by := y % 2
		println(bx, by)
		if bx != by {
			res++
		}
		x /= 2
		y /= 2
	}
	return res
}

/**
1 0
0 0
0 1
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0
0 0


得出（32位下）
1的二进制表示为 0 0000000000 0000000000 0000000001
4的二进制表示为 0 0000000000 0000000000 0000000100

挨个比较每一位，只要是不相同就+1，最后得出不同的有几位，就是结果
*/

func HammingDistance2(x int, y int) int {
	res := 0
	z := x ^ y //2个数字异或，相同结果是0，不同结果是1
	for z > 0 {
		res += z % 2
		z /= 2
	}
	return res
}

/**
2个数字异或后，计算1的位数有几位
*/
