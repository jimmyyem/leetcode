package answers

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//3. 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var left, maxSize int
	cache := make(map[byte]int, 26) //key是字符，值是上一次出现的位置

	for right := 0; right < len(s); right++ {
		if _, ok := cache[s[right]]; ok {
			left = max(left, cache[s[right]]+1)
		}
		cache[s[right]] = right
		maxSize = max(maxSize, right-left+1)
		//fmt.Println(left, right)	//(left,right) 这是最大子串区间
	}
	return maxSize
}
