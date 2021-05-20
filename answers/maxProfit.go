package answers

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
//121. 买卖股票的最佳时机
func MaxProfit121(prices []int) int {
	//max_profit := 0

	//o(n)
	//min_price := 1<<31 - 1
	//
	////尽量找到最低的买价，最高的卖价
	//for _, price := range prices {
	//	if price < min_price {
	//		min_price = price
	//	}
	//	if price-min_price > max_profit {
	//		max_profit = price - min_price
	//	}
	//}
	//return max_profit

	// //o(n2)
	// //2层循环，挨个比较每一个买价固定的时候，与每个卖价做对比 得出最大利润值
	//for i := 0; i < len(prices); i++ {
	//	for j := i + 1; j < len(prices); j++ {
	//		if prices[j]-prices[i] > max_profit {
	//			max_profit = prices[j] - prices[i]
	//		}
	//	}
	//}
	//return max_profit

	//动态规划思路
	//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/121-mai-mai-gu-piao-de-zui-jia-shi-ji-zu-1v0i/
	dp0 := 0           //一直不买
	dp1 := - prices[0] //只买了一次
	dp2 := -1 << 31    //买了一次，卖了一次

	//从第二天（i=1）开始迭代之后的状态
	for i := 1; i < len(prices); i++ {
		dp1 = max(dp1, dp0-prices[i])
		dp2 = max(dp2, dp1+prices[i])
	}

	return max(dp0, dp2)
}

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
//122. 买卖股票的最佳时机 II
func MaxProfit122(prices []int) int {
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		//由于不限制购买次数所以每次都比较，大于0的才算交易
		maxProfit += max(0, prices[i]-prices[i-1])
	}

	return maxProfit
}

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
//123. 买卖股票的最佳时机 III
func MaxProfit123(prices []int) int {
	return prices[0]
}
