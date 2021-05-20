package answers

//https://leetcode-cn.com/problems/unique-paths/
//62. 不同路径
func UniquePaths(m, n int) int {
	/**
	arr[i][j]	表示走到 (i,j) 这个位置有 arr[i][j] 种路径
	最终的结果就是 arr[m-1][n-1]

	arr[i][j] = arr[i-1][j] + arr[i][j-1]
	*/
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	//set 1st row value is 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	//set 1st column value is 1
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	//set other values
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
