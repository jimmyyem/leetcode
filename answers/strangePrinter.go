package answers

//https://leetcode-cn.com/problems/strange-printer/
//664. 奇怪的打印机
func StrangePrinter(s string) int {
	cache := make(map[rune]int, len(s))
	for _, val := range s {
		if _, ok := cache[val]; !ok {
			cache[val] = 0
		}
		cache[val]++
	}

	return len(cache)
}
