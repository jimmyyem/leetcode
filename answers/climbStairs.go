package answers

//https://leetcode-cn.com/problems/climbing-stairs/
//70. 爬楼梯
func climbStairs(n int) int {
	first, second, res := 0, 0, 1

	for i := 1; i <= n; i++ {
		first = second
		second = res
		res = first + second
	}

	return res
}

/**
f(x) = f(x-1) + f(x-2)


f(x)
f(x-1)
f(x-2)

f(4)=5
f(3)=3
f(2)=2
f(1)=1
f(0)=0
*/
