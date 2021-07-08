package answers

//什么是动态规划（Dynamic Programming）？动态规划的意义是什么？ - 帅地的回答 - 知乎
//https://www.zhihu.com/question/23995189/answer/1094101149

//问题描述：一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
func Jump(n int) int {

	/**
	dp数组记录状态，key是代表跳key级台阶，value代表跳key级台阶有几种方法
	dp[i] 的含义为：跳上一个 i 级的台阶总共有 dp[i] 种跳法
	通过1～5级台阶的归纳发现：
	1 => 1
	2 => 2
	3 => 3
	4 => 5
	5 => 8
	总结规律如下：dp[n] = dp[n-1] + dp[n-2]
	初始条件需要手动指定
	*/

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	/**
	跳上n级台阶    有几种跳法
	*/

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}