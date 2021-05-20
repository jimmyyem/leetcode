package answers

//https://leetcode-cn.com/problems/minimum-path-sum/
//64. 最小路径和
func MinPathSum(grid [][]int) int {
	/**
	dp[i][j] 到达(i,j)位置路径上的数字总和
	dp[i][j] = dp[i-1][j] + grid[i][j]
	*/

	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])
	if n <= 0 {
		return 0
	}

	//构造dp数组
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	//dp[0][0]是与grid[0][0]相等的
	dp[0][0] = grid[0][0]

	//dp(i,j) 意思是从起始位置 到 (i,j) 经历的数字总和

	//第一行的数字总和（根据 `0,0`）是可以直接算出来的
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	//第一列的数字总和（根据 `0,0`）是可以直接算出来的
	for j := 1; j < m; j++ {
		dp[j][0] = dp[j-1][0] + grid[j][0]
	}

	//其余位置的（i,j）的数字总和填充满
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}
