package answers

import (
	"strings"
)

//https://leetcode-cn.com/problems/longest-common-prefix/
//14. 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	first := strs[0]
	for i := 1; i <= len(first); i++ {
		prefix := first[:i]
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], prefix) {
				return prefix[:i-1]
			}
		}
	}

	return first
}
