package answers

//https://leetcode-cn.com/problems/longest-palindromic-substring/
//5.最长回文子串
func LongestPalindrome(s string) string {
	var start, end = 0, 0

	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)   //aba
		l2, r2 := expandAroundCenter(s, i, i+1) //abba

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

//根据给定位置开始从中心向左右两侧扩展，搜索回文内容索引位置
func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
