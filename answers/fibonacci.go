package answers

var (
	cache = make(map[int]int)
)

//动态规划方式 时间o(n)+空间o(n)
func Fibo(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

//时间o(n)+空间o(1) 最优
func Fibo2(n int) int {
	if n < 1 {
		return n
	}

	var prev, curr int = 0, 1
	for i := 2; i <= n; i++ {
		sum := curr + prev
		prev = curr
		curr = sum
	}

	return curr
}

//栈迭代方式  数字大容易栈溢出，且运算速度慢
//cache又称为备忘录
func Fibonacci(n int, useCache bool) int {
	if useCache {
		if val, ok := cache[n]; ok {
			return val
		}
	}

	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	default:
		res := Fibonacci(n-1, useCache) + Fibonacci(n-2, useCache)
		if useCache {
			cache[n] = res
		}
		return res
	}
}

//根据给定的硬币组合和总钱数，给出能满足的方案，如果不能满足返回-1
//动态规划       [1,2,5]     11
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	maxNum := 1<<31 - 1
	answer := maxNum

	for _, coin := range coins {
		if amount-coin < 0 {
			continue
		}

		sub := CoinChange(coins, amount-coin)
		if sub == -1 {
			continue
		}

		answer = min(answer, sub+1)
	}

	if answer == maxNum {
		return -1
	}
	return answer
}

func AllSort(nums []int, idx int) {
	if len(nums) > idx-1 {
		return
	}

}
