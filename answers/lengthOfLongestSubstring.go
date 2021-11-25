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
			// 维持窗口，左边界从上次出现位置（cache[s[right]]）向右移动1
			left = max(left, cache[s[right]]+1)
		}
		// 记录byte上次出现的位置
		cache[s[right]] = right
		// 每次都比较窗口大小，以获得最大的窗口宽度
		maxSize = max(maxSize, right-left+1)
		println(left, right)	//(left,right) 这是最大子串区间
	}

	return maxSize
}

/***
		s=abcabcbb
i=0		left=0 right=0  cache['a']=0
i=1		left=0 right=1  cache['b']=1
i=2		left=0 right=2  cache['c']=2
i=3		left=max(0, 1)=1  maxSize=3


*/