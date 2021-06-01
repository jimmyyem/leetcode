package answers

//https://leetcode-cn.com/problems/implement-strstr/
//28. 实现 strStr()
//Brute Force
func StrStr(haystack string, needle string) int {
	var len1, len2 = len(haystack), len(needle)

	if len2 == 0 || haystack == needle {
		return 0
	}
	if len2 > len1 {
		return -1
	}

	var i, j = 0, 0
	for i = 0; i <= len1-len2; i++ {
		for j = 0; j < len2; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len2 {
			return i
		}
	}

	return -1
}