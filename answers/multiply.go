package answers

//https://leetcode-cn.com/problems/recursive-mulitply-lcci/solution/ci-shu-zui-shao-de-di-gui-diao-yong-by-wang-ming-j/
//面试题 08.05. 递归乘法
func Multiply(a, b int) int {
	if b == 0 || a == 0 {
		return 0
	}
	if b == 1 {
		return a
	}

	var res int

	for i := 0; i < b; i++ {
		res += a
	}

	return res
}
