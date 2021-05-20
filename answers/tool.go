package answers

//工具/公用方法
func Max(x, y int) int {
	return max(x, y)
}

func Min(x, y int) int {
	return min(x, y)
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
